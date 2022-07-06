package typ

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type AnyOfRuler interface {
	Discriminator() string
	SupportedValuesByDiscriminatorValue() map[string]any
}

type AnyOfSupporter interface {
	supportAnyOf()
	Ruler() AnyOfRuler
}

func NewAnyOf[T AnyOfRuler](value any) AnyOf[T] {
	return AnyOf[T]{content: value}
}

type AnyOf[T AnyOfRuler] struct {
	ruler   T
	content any
}

func (t AnyOf[T]) supportAnyOf() {}

func (t AnyOf[T]) Ruler() AnyOfRuler {
	return t.ruler
}

func (t AnyOf[T]) Value() any {
	return t.content
}

func (t *AnyOf[T]) UnmarshalJSON(data []byte) error {
	discriminatorName := t.ruler.Discriminator()
	rawData := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &rawData)
	if err != nil {
		return fmt.Errorf("couldd not unmarshal data: %w", err)
	}
	discriminatorRawValue, ok := rawData[discriminatorName]
	if !ok {
		return fmt.Errorf(`discriminator "%s" not found`, discriminatorName)
	}
	var discriminatorValue string
	err = json.Unmarshal(discriminatorRawValue, &discriminatorValue)
	if err != nil {
		return fmt.Errorf(`discriminator value "%s" is not a string`, string(discriminatorRawValue))
	}
	concreteValueType, ok := t.ruler.SupportedValuesByDiscriminatorValue()[discriminatorValue]
	if !ok {
		return fmt.Errorf(
			`could not find value for discriminator "%s" with value %s`,
			discriminatorName,
			discriminatorValue,
		)
	}
	concreteValueTypeKind := reflect.TypeOf(concreteValueType).Kind()
	if reflect.TypeOf(concreteValueType).Kind() != reflect.Ptr {
		return fmt.Errorf(`discriminator value %s is not a pointer`, concreteValueTypeKind)
	}
	err = json.Unmarshal(data, &concreteValueType)
	if err != nil {
		return fmt.Errorf(`could not unmarshal concrete type: %w`, err)
	}
	t.content = concreteValueType
	return nil
}

func (t *AnyOf[T]) findDiscriminatorValueBySupportedValue(val any) (string, bool) {
	for typeStr, supportedVal := range t.ruler.SupportedValuesByDiscriminatorValue() {
		extractedVal := extractUnderlyingValueFromPossiblePointer(val)
		extractedSupportedVal := extractUnderlyingValueFromPossiblePointer(supportedVal)
		if reflect.TypeOf(extractedVal) == reflect.TypeOf(extractedSupportedVal) {
			return typeStr, true
		}
	}
	return "", false
}

func (t AnyOf[T]) MarshalJSON() ([]byte, error) {
	discriminatorValue, ok := t.findDiscriminatorValueBySupportedValue(t.content)
	if !ok {
		return nil, fmt.Errorf("could not find discriminator value for type: %s", reflect.TypeOf(t.content).Name())
	}
	data, err := json.Marshal(t.content)
	if err != nil {
		return nil, err
	}
	rawData := make(map[string]json.RawMessage)
	err = json.Unmarshal(data, &rawData)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal data: %w", err)
	}
	marshalledDiscriminatorValue, err := json.Marshal(discriminatorValue)
	if err != nil {
		return nil, fmt.Errorf(`could not uarshal discriminatorValue "%s"`, discriminatorValue)
	}
	rawData[t.ruler.Discriminator()] = marshalledDiscriminatorValue
	return json.Marshal(rawData)
}
