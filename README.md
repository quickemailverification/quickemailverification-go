<p align="center">
  <img src="https://quickemailverification.com/extra/images/logo_github.png" alt="Quick Email Verification">
  <br>
</p>


# Official Email Validation API library client for Go

QuickEmailVerification provides the quickest way to avoid fake / invalid emails.

* Get actual users by scanning email address.
* Remove invalid, dead and fake emails from your email list.
* Save time and money by sending mail to actual users.

## Let's Get Started

To begin, signUp at [quickemailverification.com](https://quickemailverification.com) and create a FREE account. After signup logged in, click on **[API Settings](https://quickemailverification.com/apisettings)** and then click **Add API Key**. To use API you need 2 parameters.

email - (string) This is a email address you need to verify. It should be url encoded.

apikey - (string) This is the key you generated from "api settings" page.

NOTE: Keep API key with you. You'll need it to setup the client as explained below.

## Installation

Install the package with the following command:

```bash
go get github.com/quickemailverification/quickemailverification-go
```

Import the package with the following line:

```
import "github.com/quickemailverification/quickemailverification-go"
```

## Usage

```
import "github.com/quickemailverification/quickemailverification-go"

// Create a client with your API key
qev := quickemailverification.CreateClient("YOUR_API_KEY")

// PRODUCTION MODE
// Verify an email address
response, err := qev.Verify("test@example.com")

// SANDBOX MODE
// response, err := qev.Sandbox("valid@example.com")

if err != nil {
	panic(err)
}

// The response is in the body attribute
print(response.Result)
```

### Response information

A successful API call responds with the following values:

- **result** `string` - The verified results will be: `valid`, `invalid`, `unknown`
- **reason** `string` - Reason definitions are as below:
  - `invalid_email` - Specified email has invalid email address syntax
  - `invalid_domain` - Domain name does not exist
  - `rejected_email` - SMTP server rejected email. Email does not exist
  - `accepted_email` - SMTP server accepted email address
  - `no_connect` - SMTP server connection failure
  - `timeout` -  Session time out occurred at SMTP server
  - `unavailable_smtp` - SMTP server is not available to process request
  - `unexpected_error` - An unexpected error has occurred
  - `no_mx_record` - Could not get MX records for domain
  - `temporarily_blocked` - Email is temporarily greylisted 
  - `exceeded_storage` - SMTP server rejected email. Exceeded storage allocation

- **disposable**  `true | false` - *true* if the email address uses a *disposable* domain
- **accept_all**  `true | false` - *true* if the domain appears to *accept all* emails delivered to that domain
- **role**  `true | false` - *true* if the email address is a *role* address (`manager@example.com`, `ceo@example.com`, etc)
- **free** `true | false` - *true* if the email address is from free email provider like Gmail, Yahoo!, Hotmail etc.
- **email** `string` - Returns a normalized version. (`Niki@example.com` -> `niki@example.com`)
- **user** `string` - The local part of an email address. (`niki@example.com` -> `niki`)
- **domain** `string` - The domain of the provided email address. (`niki@example.com` -> `example.com`)
- **safe_to_send** `true | false` - *true* if the email address is safe for deliverability
- **did_you_mean** `string` - Returns email suggestions if specific typo errors found in email
- **success** `true | false` - *true* if the API request was successful
- **message** `string` - Describes API call failure reason

### HTTP Response headers

Total remaining credits can be found by http response header. It contains overall remaining credits, including Persistent & Per day credits.

- **X-QEV-Remaining-Credits** - Your remaining email verification credits (i.e. Per Day Credits + Persistent Credits).

## HTTP status codes for QuickEmailVerification API calls

QuickEmailVerification API also returns following HTTP status codes to indicate success or failure of request.

- `200` - Request is completed successfully.
- `400` - Server can not understand the request sent to it. This is kind of response can occur if parameters are passed wrongly.
- `401` - Server can not verify your authentication to use api. Please check whether API key is proper or not.
- `402` - You are running out of your credit limit.
- `404` - Requested API can not be found on server.
- `429` - Too many requests. Rate limit exceeded.

## Sandbox Mode
QuickEmailVerification single email verification API sandbox mode helps developers to test their integration against simulated results. Requesting against sandbox endpoint is totally free and no credits will be deducted from actual credit quota.

Please refer our [knowledge base](http://docs.quickemailverification.com/email-verification-api/sandbox-mode) to learn more about sandbox mode.

## License
MIT

## Bug Reports
Report [here](https://github.com/quickemailverification/quickemailverification-go/issues).

## Need Help? Feel free to contact us
https://quickemailverification.com/contact-us
