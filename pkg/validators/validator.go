// http://godoc.org/gopkg.in/go-playground/validator.v9
// some common validators

// required
// len
// max
// min
// eq
// ne
// gt
// gte
// lt
// lte
// eqfield
// alpha
// alphanum
// alphaunicode
// alphanumunicode
// numeric
// hexadecimal
// hexcolor
// rgb
// rgba
// hsl
// hsla
// email
// url
// uri
// base64
// isbn
// isbn10
// isbn13
// uuid
// uuid3
// uuid4
// uuid5
// ascii
// printascii
// multybyte
// datauri
// latitude
// longtitude
// ssn
// ip
// ipv4
// ipv6
// cidrv4
// cidrv6
// tcp_addr
// tcp4_addr
// tcp6_addr
// udp_addr
// udp4_addr
// udp6_addr
// ip4_addr
// ip6_addr
// unix_addr
// mac

package validators

import (
	"errors"
	"fmt"
	"go-clean-architecture/pkg/log"
	"regexp"
	"strconv"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	logger = log.GetLogger()
	validateInstance *validator.Validate

	regex = make(map[string]*regexp.Regexp)

	aliases = map[string]string{
		"v_integer":         "omitempty,min=1",
		"v_string_isprint":  "omitempty,min=0,printascii",
		"v_string":          "omitempty,alphanumunicode",
		"v_url":             "omitempty,url",
		"v_phone":           "omitempty,alphanum,min=10,max=11",
		"v_email":           "omitempty,email",
		"v_bool":            "omitempty,eq=0|eq=1|eq=t|eq=f",
		"v_unit":            "omitempty,eq=vnd|eq=credit",
	}

	customValidateFuncs = map[string]validator.Func{
		"region":                 validateRegion,
		"name":                   validateName,
		"unicode":                validateUnicode,
	}
)

func vRegex(name string, value string) bool {
	return regex[name].MatchString(value)
}

func validateSMSID(fl validator.FieldLevel) bool {
	smsID := fl.Field().String()
	paymentType := strings.ToLower(fl.Parent().FieldByName("PaymentType").String())
	return paymentType != "sms" || smsID == ""
}

func validateSMSProviderReference(fl validator.FieldLevel) bool {
	smsProviderRef := fl.Field().String()
	paymentType := strings.ToLower(fl.Parent().FieldByName("PaymentType").String())
	return paymentType != "sms" || smsProviderRef == ""
}

func validatePrivs(fl validator.FieldLevel) bool {
	return vRegex("v_privs", strings.ToLower(fl.Field().String()))
}

func validateRegion(fl validator.FieldLevel) bool {
	return vRegex("v_region", strconv.FormatInt(fl.Field().Int(), 10))
}

func validateName(fl validator.FieldLevel) bool {
	return vRegex("v_name", strings.ToLower(fl.Field().String()))
}

func validateUnicode(fl validator.FieldLevel) bool {
	return vRegex("v_unicode", strings.ToLower(fl.Field().String()))
}

func validateUID(fl validator.FieldLevel) bool {
	return vRegex("v_uid", strconv.FormatInt(fl.Field().Int(), 10))
}

func validateGiveTrade(fl validator.FieldLevel) bool {
	return vRegex("v_givetrade", strings.ToLower(fl.Field().String()))
}

func validateProductDate(fl validator.FieldLevel) bool {
	return vRegex("v_product_date", strings.ToLower(fl.Field().String()))
}

func validateHost(fl validator.FieldLevel) bool {
	return vRegex("v_host", strings.ToLower(fl.Field().String()))
}

func validateAppl(fl validator.FieldLevel) bool {
	return vRegex("v_appl", strings.ToLower(fl.Field().String()))
}

func validateSellerAddr(fl validator.FieldLevel) bool {
	return vRegex("v_seller_addr", strings.ToLower(fl.Field().String()))
}

func validatePaymentType(fl validator.FieldLevel) bool {
	return vRegex("v_pay_payment_type", strings.ToLower(fl.Field().String()))
}

func validateServiceType(fl validator.FieldLevel) bool {
	return vRegex("v_service_type", strings.ToLower(fl.Field().String()))
}

func validatePaymentAction(fl validator.FieldLevel) bool {
	return vRegex("v_pay_payment_action", strings.ToLower(fl.Field().String()))
}

func validateServiceAction(fl validator.FieldLevel) bool {
	return vRegex("v_service_action", strings.ToLower(fl.Field().String()))
}

func validateServiceNoActionType(fl validator.FieldLevel) bool {
	return vRegex("service_no_action_type", strings.ToLower(fl.Field().String()))
}

func validateBconfKey(fl validator.FieldLevel) bool {
	fval := fl.Field().String()
	return len(fval) > 0 && len(fval) < 8193 && vRegex("v_bconf_keyval", fval)
}

func initRegex() {
	patterns := map[string]string{
		"v_region":             `^[1-3]?[0-9](,[1-3]?[0-9])*$`,
		"v_name":               `^[\x20-\x{23CD}\x{2E80}-\x{4DFF}\x{4E00}-\x{9FFF}\x{F900}-\x{FAFF}\x{FF00}-\x{FF5A}\x{1BCA0}-\x{2A6DF}\x{2A700}-\x{2B73F}\x{2F800}-\x{2FA1F}]+`,
		"v_unicode":            `^[\x20-\x{23CD}\x{2E80}-\x{4DFF}\x{4E00}-\x{9FFF}\x{F900}-\x{FAFF}\x{FF00}-\x{FF5A}\x{1BCA0}-\x{2A6DF}\x{2A700}-\x{2B73F}\x{2F800}-\x{2FA1F}]+`,
	}

	// Compile and panic on errors
	for name, pattern := range patterns {
		regex[name] = regexp.MustCompile(pattern)
	}
}

func init() {
	validateInstance = validator.New()

	initRegex()
	for k, v := range customValidateFuncs {
		validateInstance.RegisterValidation(k, v)
		validateInstance.RegisterAlias("v_"+k, "omitempty,"+k)
	}

	for k, v := range aliases {
		validateInstance.RegisterAlias(k, v)
	}
}

const (
	fieldErrMsg         = "%s không hợp lệ, vui lòng chỉnh sửa lại."
	defaultTranslateVal = "Thông tin"
)

var fieldTranslate = map[string]string{}

func translateField(in string) string {
	lIn := strings.ToLower(in)
	if tmp := fieldTranslate[lIn]; tmp != "" {
		return tmp
	}

	return lIn
}

// Validate the input
func Validate(i interface{}) error {
	if validateInstance == nil {
		return errors.New("PANIC: VALIDATOR NOT INIT")
	}

	err := validateInstance.Struct(i)
	switch err.(type) {
	case *validator.InvalidValidationError:
	case validator.ValidationErrors:
		v := err.(validator.ValidationErrors)[0]
		err = fmt.Errorf(fmt.Sprintf(fieldErrMsg, translateField(v.StructField())))
	}

	if err != nil {
		logger.Errorf("validate err: %v", err)
	}

	return err
}
