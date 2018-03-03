package gtranslate

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestTranslateRuUk(t *testing.T) {
    result, _ := Translate("стол", "ru", "uk")

    assert.Equal(t, result, "стіл")
}

func TestTranslateUkRu(t *testing.T) {
    result, _ := Translate("стіл", "uk", "ru")

    assert.Equal(t, result, "стол")
}

func TestTranslateRuEn(t *testing.T) {
    result, _ := Translate("стол", "ru", "en")

    assert.Equal(t, result, "table")
}

func TestTranslateUaEn(t *testing.T) {
    result, _ := Translate("стіл", "uk", "en")

    assert.Equal(t, result, "table")
}

func TestTranslateEnUa(t *testing.T) {
    result, _ := Translate("table", "en", "uk")

    assert.Equal(t, result, "стіл")
}

func TestTranslateRuUkWithQuotes(t *testing.T) {
    result, _ := Translate(`"Стол"`, "ru", "uk")

    assert.Equal(t, result, `\"Стіл\"`)
}
