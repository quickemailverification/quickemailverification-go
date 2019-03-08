package quickemailverification

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestVerifyValidEmail(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("valid@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "valid")
    assert.Equal(t, response.Reason, "accepted_email")
    assert.Equal(t, response.Disposable, "true")
    assert.Equal(t, response.AcceptAll, "true")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "valid@example.com")
    assert.Equal(t, response.User, "valid")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifySafeToSendEmail(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("safe-to-send@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "valid")
    assert.Equal(t, response.Reason, "accepted_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "safe-to-send@example.com")
    assert.Equal(t, response.User, "safe-to-send")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "true")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyFreeEmail(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("free@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "valid")
    assert.Equal(t, response.Reason, "accepted_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "true")
    assert.Equal(t, response.Email, "free@example.com")
    assert.Equal(t, response.User, "free")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "true")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyRejectedEmail(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("rejected-email@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "invalid")
    assert.Equal(t, response.Reason, "rejected_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "rejected-email@example.com")
    assert.Equal(t, response.User, "rejected-email")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyInvalidDomain(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("invalid-domain@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "invalid")
    assert.Equal(t, response.Reason, "invalid_domain")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "invalid-domain@example.com")
    assert.Equal(t, response.User, "invalid-domain")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyInvalidEmail(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("invalid-email@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "invalid")
    assert.Equal(t, response.Reason, "invalid_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "invalid-email@example.com")
    assert.Equal(t, response.User, "invalid-email")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyExceededStorage(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("exceeded-storage@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "invalid")
    assert.Equal(t, response.Reason, "exceeded_storage")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "exceeded-storage@example.com")
    assert.Equal(t, response.User, "exceeded-storage")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyNoMXRecord(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("no-mx-record@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "invalid")
    assert.Equal(t, response.Reason, "no_mx_record")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "no-mx-record@example.com")
    assert.Equal(t, response.User, "no-mx-record")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyDidYouMean(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("did-you-mean@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "invalid")
    assert.Equal(t, response.Reason, "rejected_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "did-you-mean@example.com")
    assert.Equal(t, response.User, "did-you-mean")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "did-you-mean@example.com")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyTimeout(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("timeout@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "unknown")
    assert.Equal(t, response.Reason, "timeout")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "timeout@example.com")
    assert.Equal(t, response.User, "timeout")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyUnexpectedError(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("unexpected-error@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "unknown")
    assert.Equal(t, response.Reason, "unexpected_error")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "unexpected-error@example.com")
    assert.Equal(t, response.User, "unexpected-error")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyNoConnect(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("no-connect@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "unknown")
    assert.Equal(t, response.Reason, "no_connect")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "no-connect@example.com")
    assert.Equal(t, response.User, "no-connect")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyUnavailableSMTP(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("unavailable-smtp@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "unknown")
    assert.Equal(t, response.Reason, "unavailable_smtp")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "unavailable-smtp@example.com")
    assert.Equal(t, response.User, "unavailable-smtp")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyTemporarilyBlocked(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("temporarily-blocked@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "unknown")
    assert.Equal(t, response.Reason, "temporarily_blocked")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "temporarily-blocked@example.com")
    assert.Equal(t, response.User, "temporarily-blocked")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyAcceptAll(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("accept-all@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "valid")
    assert.Equal(t, response.Reason, "accepted_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "true")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "accept-all@example.com")
    assert.Equal(t, response.User, "accept-all")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyRole(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("role@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "valid")
    assert.Equal(t, response.Reason, "accepted_email")
    assert.Equal(t, response.Disposable, "false")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "true")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "role@example.com")
    assert.Equal(t, response.User, "role")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyDisposable(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("disposable@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Result, "valid")
    assert.Equal(t, response.Reason, "accepted_email")
    assert.Equal(t, response.Disposable, "true")
    assert.Equal(t, response.AcceptAll, "false")
    assert.Equal(t, response.Role, "false")
    assert.Equal(t, response.Free, "false")
    assert.Equal(t, response.Email, "disposable@example.com")
    assert.Equal(t, response.User, "disposable")
    assert.Equal(t, response.Domain, "example.com")
    assert.Equal(t, response.MxRecord, "")
    assert.Equal(t, response.MxDomain, "")
    assert.Equal(t, response.SafeToSend, "false")
    assert.Equal(t, response.Suggested, "")
    assert.Equal(t, response.Success, "true")
}

func TestVerifyLowCredit(t *testing.T) {

    client := CreateClient(os.Getenv("apikey"))
    response, err := client.Sandbox("low-credit@example.com")

    if err != nil {
        panic(err)
    }

    assert.Equal(t, response.Success, "false")
    assert.Equal(t, response.Message, "Low credit. Payment required")
}
