package entity

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntityChecker struct{}

func (h EntityChecker) Check(entityDefinition EntityDefinition, entity Entity) error {
	// check the type of entity attributes
	for name, attribute := range entity.Attributes {
		// get the type in definition
		attributeType := entityDefinition.Attributes[name].Type
		if attributeType == "" {
			return errors.New("Missing type attribute for " + name)
		}
		err := h.checkAttributeType(name, attributeType, attribute)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h EntityChecker) checkAttributeType(attributeName string, attributeTypeDefinition Type, attributeValue interface{}) error {
	valueType := reflect.TypeOf(attributeValue)
	switch attributeTypeDefinition {
	case String:
		{
			if valueType.Kind() != reflect.String {
				return errors.New(attributeName + " isn't a String")
			} else {
				return nil
			}
		}
	case Integer:
		{
			if valueType.Kind() != reflect.Int && valueType.Kind() != reflect.Int8 && valueType.Kind() != reflect.Int16 && valueType.Kind() != reflect.Int32 && valueType.Kind() != reflect.Int64 && valueType.Kind() != reflect.Uint && valueType.Kind() != reflect.Uint8 && valueType.Kind() != reflect.Uint16 && valueType.Kind() != reflect.Uint32 && valueType.Kind() != reflect.Uint64 {
				return errors.New(attributeName + " isn't an Integer")
			} else {
				return nil
			}
		}
	case Decimal:
		{
			if valueType.Kind() != reflect.Float32 && valueType.Kind() != reflect.Float64 {
				return errors.New(attributeName + " isn't a Decimal")
			} else {
				return nil
			}
		}
	case Boolean:
		{
			if valueType.Kind() != reflect.Bool {
				return errors.New(attributeName + " isn't a Boolean")
			} else {
				return nil
			}
		}
	case ObjectId:
		{
			if valueType.Kind() == reflect.String {
				_, err := primitive.ObjectIDFromHex(attributeValue.(string))
				if err != nil {
					return errors.New(attributeName + " isn't an ObjectId")
				}
				return nil
			} else {
				return errors.New(attributeName + " isn't an ObjectId")
			}
		}
	default:
		return errors.New("type not supported")
	}
}
