/*
SendPost API

# Introduction  SendPost provides email API and SMTP relay which can be used not just to send & measure but also alert & optimised email sending.  You can use SendPost to:  * Send personalised emails to multiple recipients using email API   * Track opens and clicks  * Analyse statistics around open, clicks, bounce, unsubscribe and spam    At and advanced level you can use it to:  * Manage multiple sub-accounts which may map to your promotional or transactional sending, multiple product lines or multiple customers   * Classify your emails using groups for better analysis  * Analyse and fix email sending at sub-account level, IP Pool level or group level  * Have automated alerts to notify disruptions regarding email sending  * Manage different dedicated IP Pools so to better control your email sending  * Automatically know when IP or domain is blacklisted or sender score is down  * Leverage pro deliverability tools to get significantly better email deliverability & inboxing   [<img src=\"https://run.pstmn.io/button.svg\" alt=\"Run In Postman\" style=\"width: 128px; height: 32px;\">](https://god.gw.postman.com/run-collection/33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b%26entityType%3Dcollection%26workspaceId%3D6b1e4f65-96a9-4136-9512-6266c852517e)   # Overview  ## REST API  SendPost API is built on REST API principles. Authenticated users can interact with any of the API endpoints to perform:  * **GET**- to get a resource  * **POST** - to create a resource  * **PUT** - to update an existing resource  * **DELETE** - to delete a resource   The API endpoint for all API calls is: <code>https://api.sendpost.io/api/v1</code>   Some conventions that have been followed in the API design overall are following:   * All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header.  * All resource endpoints end with singular name and not plural. So we have <code>domain</code> instead of domains for domain resource endpoint. Likewise we have <code>sender</code> instead of senders for sender resource endpoint.  * Body submitted for POST / PUT API calls as well as JSON response from SendPost API follow camelcase convention  * All timestamps returned in response (created or submittedAt response fields) are UNIX nano epoch timestamp.   <aside class=\"success\"> All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header. </aside>   SendPost uses conventional HTTP response codes to indicate the success or failure of an API request.    * Codes in the <code>2xx</code> range indicate success.   * Codes in the <code>4xx</code> range indicate an error owing due to unauthorize access, incorrect request parameters or body etc.  * Code in the <code>5xx</code> range indicate an eror with SendPost's servers ( internal service issue or maintenance )   <aside class=\"info\"> SendPost all responses return <code>created</code> in UNIX nano epoch timestamp.  </aside>   ## Authentication  SendPost uses API keys for authentication. You can register a new SendPost API key at our [developer portal](https://app.sendpost.io/register).   SendPost expects the API key to be included in all API requests to the server in a header that looks like the following:   `X-SubAccount-ApiKey: AHEZEP8192SEGH`   This API key is used for all Sub-Account level operations such as:  * Sending emails  * Retrieving stats regarding open, click, bounce, unsubscribe and spam  * Uploading suppressions list  * Verifying sending domains and more  In addition to <code>X-SubAccount-ApiKey</code> you also have another API Key <code>X-Account-APIKey</code> which is used for Account level operations such as :  * Creating and managing sub-accounts  * Allocating IPs for your account  * Getting overall billing and usage information  * Email List validation  * Creating and managing alerts and more   <aside class=\"notice\"> You must look at individual API reference page to look at whether <code>X-SubAccount-ApiKey</code> is required or <code>X-Account-ApiKey</code> </aside>   In case an incorrect API Key header is specified or if it is missed you will get HTTP Response 401 ( Unauthorized ) response from SendPost.   ## HTTP Response Headers   Code           | Reason                 | Details ---------------| -----------------------| ----------- 200            | Success                | Everything went well 401            | Unauthorized           | Incorrect or missing API header either <code>X-SubAccount-ApiKey</code> or <code>X-Account-ApiKey</code> 403            | Forbidden              | Typically sent when resource with same name or details already exist 406            | Missing resource id    | Resource id specified is either missing or doesn't exist 422            | Unprocessable entity   | Request body is not in proper format 500            | Internal server error  | Some error happened at SendPost while processing API request 503            | Service Unavailable    | SendPost is offline for maintenance. Please try again later  # API SDKs  We have native SendPost SDKs in the following programming languages. You can integrate with them or create your own SDK with our API specification. In case you need any assistance with respect to API then do reachout to our team from website chat or email us at **hello@sendpost.io**   * [PHP](https://github.com/sendpost/sendpost_php_sdk)  * [Javascript](https://github.com/sendpost/sendpost_javascript_sdk)  * [Ruby](https://github.com/sendpost/sendpost_ruby_sdk)  * [Python](https://github.com/sendpost/sendpost_python_sdk)  * [Golang](https://github.com/sendpost/sendpost_go_sdk)   # API Reference  SendX REST API can be broken down into two major sub-sections:   * Sub-Account  * Account    Sub-Account API operations enable common email sending API use-cases like sending bulk email, adding new domains or senders for email sending programmatically, retrieving stats, adding suppressions etc. All Sub-Account API operations need to pass <code>X-SubAccount-ApiKey</code> header with every API call.   The Account API operations allow users to manage multiple sub-accounts and manage IPs. A single parent SendPost account can have 100's of sub-accounts. You may want to create sub-accounts for different products your company is running or to segregate types of emails or for managing email sending across multiple customers of yours.   # SMTP Reference  Simple Mail Transfer Protocol (SMTP) is a quick and easy way to send email from one server to another. SendPost provides an SMTP service that allows you to deliver your email via our servers instead of your own client or server.  This means you can count on SendPost's delivery at scale for your SMTP needs.    ## Integrating SMTP    1. Get the SMTP `username` and `password` from your SendPost account.  2. Set the server host in your email client or application to `smtp.sendpost.io`. This setting is sometimes referred to as the external SMTP server or the SMTP relay.  3. Set the `username` and `password`.  4. Set the port to `587` (or as specified below).  ## SMTP Ports   - For an unencrypted or a TLS connection, use port `25`, `2525` or `587`.  - For a SSL connection, use port `465`  - Check your firewall and network to ensure they're not blocking any of our SMTP Endpoints.   SendPost supports STARTTLS for establishing a TLS-encrypted connection. STARTTLS is a means of upgrading an unencrypted connection to an encrypted connection. There are versions of STARTTLS for a variety of protocols; the SMTP version is defined in [RFC 3207](https://www.ietf.org/rfc/rfc3207.txt).   To set up a STARTTLS connection, the SMTP client connects to the SendPost SMTP endpoint `smtp.sendpost.io` on port 25, 587, or 2525, issues an EHLO command, and waits for the server to announce that it supports the STARTTLS SMTP extension. The client then issues the STARTTLS command, initiating TLS negotiation. When negotiation is complete, the client issues an EHLO command over the new encrypted connection, and the SMTP session proceeds normally.   <aside class=\"success\"> If you are unsure which port to use, a TLS connection on port 587 is typically recommended. </aside>   ## Sending email from your application   ```javascript \"use strict\";  const nodemailer = require(\"nodemailer\");  async function main() { // create reusable transporter object using the default SMTP transport let transporter = nodemailer.createTransport({ host: \"smtp.sendpost.io\", port: 587, secure: false, // true for 465, false for other ports auth: { user:  \"<username>\" , // generated ethereal user pass: \"<password>\", // generated ethereal password }, requireTLS: true, debug: true, logger: true, });  // send mail with defined transport object try { let info = await transporter.sendMail({ from: 'erlich@piedpiper.com', to: 'gilfoyle@piedpiper.com', subject: 'Test Email Subject', html: '<h1>Hello Geeks!!!</h1>', }); console.log(\"Message sent: %s\", info.messageId); } catch (e) { console.log(e) } }  main().catch(console.error); ```  For PHP   ```php <?php // Import PHPMailer classes into the global namespace use PHPMailer\\PHPMailer\\PHPMailer; use PHPMailer\\PHPMailer\\SMTP; use PHPMailer\\PHPMailer\\Exception;  // Load Composer's autoloader require 'vendor/autoload.php';  $mail = new PHPMailer(true);  // Settings try { $mail->SMTPDebug = SMTP::DEBUG_CONNECTION;                  // Enable verbose debug output $mail->isSMTP();                                            // Send using SMTP $mail->Host       = 'smtp.sendpost.io';                     // Set the SMTP server to send through $mail->SMTPAuth   = true;                                   // Enable SMTP authentication $mail->Username   = '<username>';                           // SMTP username $mail->Password   = '<password>';                           // SMTP password $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;         // Enable implicit TLS encryption $mail->Port       = 587;                                    // TCP port to connect to; use 587 if you have set `SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS`  //Recipients $mail->setFrom('erlich@piedpiper.com', 'Erlich'); $mail->addAddress('gilfoyle@piedpiper.com', 'Gilfoyle');  //Content $mail->isHTML(true);                                  //Set email format to HTML $mail->Subject = 'Here is the subject'; $mail->Body    = 'This is the HTML message body <b>in bold!</b>'; $mail->AltBody = 'This is the body in plain text for non-HTML mail clients';  $mail->send(); echo 'Message has been sent';  } catch (Exception $e) { echo \"Message could not be sent. Mailer Error: {$mail->ErrorInfo}\"; } ``` For Python ```python #!/usr/bin/python3  import sys import os import re  from smtplib import SMTP import ssl  from email.mime.text import MIMEText  SMTPserver = 'smtp.sendpost.io' PORT = 587 sender =     'erlich@piedpiper.com' destination = ['gilfoyle@piedpiper.com']  USERNAME = \"<username>\" PASSWORD = \"<password>\"  # typical values for text_subtype are plain, html, xml text_subtype = 'plain'  content=\"\"\"\\ Test message \"\"\"  subject=\"Sent from Python\"  try: msg = MIMEText(content, text_subtype) msg['Subject']= subject msg['From']   = sender  conn = SMTP(SMTPserver, PORT) conn.ehlo() context = ssl.create_default_context() conn.starttls(context=context)  # upgrade to tls conn.ehlo() conn.set_debuglevel(True) conn.login(USERNAME, PASSWORD)  try: resp = conn.sendmail(sender, destination, msg.as_string()) print(\"Send Mail Response: \", resp) except Exception as e: print(\"Send Email Error: \", e) finally: conn.quit()  except Exception as e: print(\"Error:\", e) ``` For Golang ```go package main  import ( \"fmt\" \"net/smtp\" \"os\" )  // Sending Email Using Smtp in Golang  func main() {  username := \"<username>\" password := \"<password>\"  from := \"erlich@piedpiper.com\" toList := []string{\"gilfoyle@piedpiper.com\"} host := \"smtp.sendpost.io\" port := \"587\" // recommended  // This is the message to send in the mail msg := \"Hello geeks!!!\"  // We can't send strings directly in mail, // strings need to be converted into slice bytes body := []byte(msg)  // PlainAuth uses the given username and password to // authenticate to host and act as identity. // Usually identity should be the empty string, // to act as username. auth := smtp.PlainAuth(\"\", username, password, host)  // SendMail uses TLS connection to send the mail // The email is sent to all address in the toList, // the body should be of type bytes, not strings // This returns error if any occured. err := smtp.SendMail(host+\":\"+port, auth, from, toList, body)  // handling the errors if err != nil { fmt.Println(err) os.Exit(1) }  fmt.Println(\"Successfully sent mail to all user in toList\") }  ``` For Java ```java // implementation 'com.sun.mail:javax.mail:1.6.2'  import java.util.Properties;  import javax.mail.Message; import javax.mail.Session; import javax.mail.Transport; import javax.mail.internet.InternetAddress; import javax.mail.internet.MimeMessage;  public class SMTPConnect {  // This address must be verified. static final String FROM = \"erlich@piedpiper.com\"; static final String FROMNAME = \"Erlich Bachman\";  // Replace recipient@example.com with a \"To\" address. If your account // is still in the sandbox, this address must be verified. static final String TO = \"gilfoyle@piedpiper.com\";  // Replace smtp_username with your SendPost SMTP user name. static final String SMTP_USERNAME = \"<username>\";  // Replace smtp_password with your SendPost SMTP password. static final String SMTP_PASSWORD = \"<password>\";  // SMTP Host Name static final String HOST = \"smtp.sendpost.io\";  // The port you will connect to on SendPost SMTP Endpoint. static final int PORT = 587;  static final String SUBJECT = \"SendPost SMTP Test (SMTP interface accessed using Java)\";  static final String BODY = String.join( System.getProperty(\"line.separator\"), \"<h1>SendPost SMTP Test</h1>\", \"<p>This email was sent with SendPost using the \", \"<a href='https://github.com/eclipse-ee4j/mail'>Javamail Package</a>\", \" for <a href='https://www.java.com'>Java</a>.\" );  public static void main(String[] args) throws Exception {  // Create a Properties object to contain connection configuration information. Properties props = System.getProperties(); props.put(\"mail.transport.protocol\", \"smtp\"); props.put(\"mail.smtp.port\", PORT); props.put(\"mail.smtp.starttls.enable\", \"true\"); props.put(\"mail.smtp.debug\", \"true\"); props.put(\"mail.smtp.auth\", \"true\");  // Create a Session object to represent a mail session with the specified properties. Session session = Session.getDefaultInstance(props);  // Create a message with the specified information. MimeMessage msg = new MimeMessage(session); msg.setFrom(new InternetAddress(FROM,FROMNAME)); msg.setRecipient(Message.RecipientType.TO, new InternetAddress(TO)); msg.setSubject(SUBJECT); msg.setContent(BODY,\"text/html\");  // Create a transport. Transport transport = session.getTransport();  // Send the message. try { System.out.println(\"Sending...\");  // Connect to SendPost SMTP using the SMTP username and password you specified above. transport.connect(HOST, SMTP_USERNAME, SMTP_PASSWORD);  // Send the email. transport.sendMessage(msg, msg.getAllRecipients()); System.out.println(\"Email sent!\");  } catch (Exception ex) {  System.out.println(\"The email was not sent.\"); System.out.println(\"Error message: \" + ex.getMessage()); System.out.println(ex); } // Close and terminate the connection. } } ```  Many programming languages support sending email using SMTP. This capability might be built into the programming language itself, or it might be available as an add-on, plug-in, or library. You can take advantage of this capability by sending email through SendPost from within application programs that you write.  We have provided examples in Python3, Golang, Java, PHP, JS. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
)

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event struct for Event
type Event struct {
	EventID *string `json:"eventID,omitempty"`
	Groups []string `json:"groups,omitempty"`
	IpID *int32 `json:"ipID,omitempty"`
	IpPoolID *int32 `json:"ipPoolID,omitempty"`
	DomainID *int32 `json:"domainID,omitempty"`
	TpspId *int32 `json:"tpspId,omitempty"`
	MessageType *string `json:"messageType,omitempty"`
	MessageSubject *string `json:"messageSubject,omitempty"`
	AccountID *int32 `json:"accountID,omitempty"`
	SubAccountID *int32 `json:"subAccountID,omitempty"`
	MessageID *string `json:"messageID,omitempty"`
	Type *int32 `json:"type,omitempty"`
	From *string `json:"from,omitempty"`
	FromName *string `json:"fromName,omitempty"`
	To *string `json:"to,omitempty"`
	ToName *string `json:"toName,omitempty"`
	SubmittedAt *int32 `json:"submittedAt,omitempty"`
	SmtpCode *int32 `json:"smtpCode,omitempty"`
	SmtpDescription *string `json:"smtpDescription,omitempty"`
	EventMetadata *EventMetadata `json:"eventMetadata,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *Event) GetEventID() string {
	if o == nil || IsNil(o.EventID) {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventIDOk() (*string, bool) {
	if o == nil || IsNil(o.EventID) {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *Event) HasEventID() bool {
	if o != nil && !IsNil(o.EventID) {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *Event) SetEventID(v string) {
	o.EventID = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Event) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Event) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *Event) SetGroups(v []string) {
	o.Groups = v
}

// GetIpID returns the IpID field value if set, zero value otherwise.
func (o *Event) GetIpID() int32 {
	if o == nil || IsNil(o.IpID) {
		var ret int32
		return ret
	}
	return *o.IpID
}

// GetIpIDOk returns a tuple with the IpID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIpIDOk() (*int32, bool) {
	if o == nil || IsNil(o.IpID) {
		return nil, false
	}
	return o.IpID, true
}

// HasIpID returns a boolean if a field has been set.
func (o *Event) HasIpID() bool {
	if o != nil && !IsNil(o.IpID) {
		return true
	}

	return false
}

// SetIpID gets a reference to the given int32 and assigns it to the IpID field.
func (o *Event) SetIpID(v int32) {
	o.IpID = &v
}

// GetIpPoolID returns the IpPoolID field value if set, zero value otherwise.
func (o *Event) GetIpPoolID() int32 {
	if o == nil || IsNil(o.IpPoolID) {
		var ret int32
		return ret
	}
	return *o.IpPoolID
}

// GetIpPoolIDOk returns a tuple with the IpPoolID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIpPoolIDOk() (*int32, bool) {
	if o == nil || IsNil(o.IpPoolID) {
		return nil, false
	}
	return o.IpPoolID, true
}

// HasIpPoolID returns a boolean if a field has been set.
func (o *Event) HasIpPoolID() bool {
	if o != nil && !IsNil(o.IpPoolID) {
		return true
	}

	return false
}

// SetIpPoolID gets a reference to the given int32 and assigns it to the IpPoolID field.
func (o *Event) SetIpPoolID(v int32) {
	o.IpPoolID = &v
}

// GetDomainID returns the DomainID field value if set, zero value otherwise.
func (o *Event) GetDomainID() int32 {
	if o == nil || IsNil(o.DomainID) {
		var ret int32
		return ret
	}
	return *o.DomainID
}

// GetDomainIDOk returns a tuple with the DomainID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDomainIDOk() (*int32, bool) {
	if o == nil || IsNil(o.DomainID) {
		return nil, false
	}
	return o.DomainID, true
}

// HasDomainID returns a boolean if a field has been set.
func (o *Event) HasDomainID() bool {
	if o != nil && !IsNil(o.DomainID) {
		return true
	}

	return false
}

// SetDomainID gets a reference to the given int32 and assigns it to the DomainID field.
func (o *Event) SetDomainID(v int32) {
	o.DomainID = &v
}

// GetTpspId returns the TpspId field value if set, zero value otherwise.
func (o *Event) GetTpspId() int32 {
	if o == nil || IsNil(o.TpspId) {
		var ret int32
		return ret
	}
	return *o.TpspId
}

// GetTpspIdOk returns a tuple with the TpspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTpspIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TpspId) {
		return nil, false
	}
	return o.TpspId, true
}

// HasTpspId returns a boolean if a field has been set.
func (o *Event) HasTpspId() bool {
	if o != nil && !IsNil(o.TpspId) {
		return true
	}

	return false
}

// SetTpspId gets a reference to the given int32 and assigns it to the TpspId field.
func (o *Event) SetTpspId(v int32) {
	o.TpspId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *Event) GetMessageType() string {
	if o == nil || IsNil(o.MessageType) {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMessageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *Event) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *Event) SetMessageType(v string) {
	o.MessageType = &v
}

// GetMessageSubject returns the MessageSubject field value if set, zero value otherwise.
func (o *Event) GetMessageSubject() string {
	if o == nil || IsNil(o.MessageSubject) {
		var ret string
		return ret
	}
	return *o.MessageSubject
}

// GetMessageSubjectOk returns a tuple with the MessageSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMessageSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.MessageSubject) {
		return nil, false
	}
	return o.MessageSubject, true
}

// HasMessageSubject returns a boolean if a field has been set.
func (o *Event) HasMessageSubject() bool {
	if o != nil && !IsNil(o.MessageSubject) {
		return true
	}

	return false
}

// SetMessageSubject gets a reference to the given string and assigns it to the MessageSubject field.
func (o *Event) SetMessageSubject(v string) {
	o.MessageSubject = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *Event) GetAccountID() int32 {
	if o == nil || IsNil(o.AccountID) {
		var ret int32
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAccountIDOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountID returns a boolean if a field has been set.
func (o *Event) HasAccountID() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.
func (o *Event) SetAccountID(v int32) {
	o.AccountID = &v
}

// GetSubAccountID returns the SubAccountID field value if set, zero value otherwise.
func (o *Event) GetSubAccountID() int32 {
	if o == nil || IsNil(o.SubAccountID) {
		var ret int32
		return ret
	}
	return *o.SubAccountID
}

// GetSubAccountIDOk returns a tuple with the SubAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSubAccountIDOk() (*int32, bool) {
	if o == nil || IsNil(o.SubAccountID) {
		return nil, false
	}
	return o.SubAccountID, true
}

// HasSubAccountID returns a boolean if a field has been set.
func (o *Event) HasSubAccountID() bool {
	if o != nil && !IsNil(o.SubAccountID) {
		return true
	}

	return false
}

// SetSubAccountID gets a reference to the given int32 and assigns it to the SubAccountID field.
func (o *Event) SetSubAccountID(v int32) {
	o.SubAccountID = &v
}

// GetMessageID returns the MessageID field value if set, zero value otherwise.
func (o *Event) GetMessageID() string {
	if o == nil || IsNil(o.MessageID) {
		var ret string
		return ret
	}
	return *o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMessageIDOk() (*string, bool) {
	if o == nil || IsNil(o.MessageID) {
		return nil, false
	}
	return o.MessageID, true
}

// HasMessageID returns a boolean if a field has been set.
func (o *Event) HasMessageID() bool {
	if o != nil && !IsNil(o.MessageID) {
		return true
	}

	return false
}

// SetMessageID gets a reference to the given string and assigns it to the MessageID field.
func (o *Event) SetMessageID(v string) {
	o.MessageID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Event) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Event) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Event) SetType(v int32) {
	o.Type = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Event) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Event) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *Event) SetFrom(v string) {
	o.From = &v
}

// GetFromName returns the FromName field value if set, zero value otherwise.
func (o *Event) GetFromName() string {
	if o == nil || IsNil(o.FromName) {
		var ret string
		return ret
	}
	return *o.FromName
}

// GetFromNameOk returns a tuple with the FromName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetFromNameOk() (*string, bool) {
	if o == nil || IsNil(o.FromName) {
		return nil, false
	}
	return o.FromName, true
}

// HasFromName returns a boolean if a field has been set.
func (o *Event) HasFromName() bool {
	if o != nil && !IsNil(o.FromName) {
		return true
	}

	return false
}

// SetFromName gets a reference to the given string and assigns it to the FromName field.
func (o *Event) SetFromName(v string) {
	o.FromName = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *Event) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *Event) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *Event) SetTo(v string) {
	o.To = &v
}

// GetToName returns the ToName field value if set, zero value otherwise.
func (o *Event) GetToName() string {
	if o == nil || IsNil(o.ToName) {
		var ret string
		return ret
	}
	return *o.ToName
}

// GetToNameOk returns a tuple with the ToName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetToNameOk() (*string, bool) {
	if o == nil || IsNil(o.ToName) {
		return nil, false
	}
	return o.ToName, true
}

// HasToName returns a boolean if a field has been set.
func (o *Event) HasToName() bool {
	if o != nil && !IsNil(o.ToName) {
		return true
	}

	return false
}

// SetToName gets a reference to the given string and assigns it to the ToName field.
func (o *Event) SetToName(v string) {
	o.ToName = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *Event) GetSubmittedAt() int32 {
	if o == nil || IsNil(o.SubmittedAt) {
		var ret int32
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSubmittedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.SubmittedAt) {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *Event) HasSubmittedAt() bool {
	if o != nil && !IsNil(o.SubmittedAt) {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given int32 and assigns it to the SubmittedAt field.
func (o *Event) SetSubmittedAt(v int32) {
	o.SubmittedAt = &v
}

// GetSmtpCode returns the SmtpCode field value if set, zero value otherwise.
func (o *Event) GetSmtpCode() int32 {
	if o == nil || IsNil(o.SmtpCode) {
		var ret int32
		return ret
	}
	return *o.SmtpCode
}

// GetSmtpCodeOk returns a tuple with the SmtpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSmtpCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.SmtpCode) {
		return nil, false
	}
	return o.SmtpCode, true
}

// HasSmtpCode returns a boolean if a field has been set.
func (o *Event) HasSmtpCode() bool {
	if o != nil && !IsNil(o.SmtpCode) {
		return true
	}

	return false
}

// SetSmtpCode gets a reference to the given int32 and assigns it to the SmtpCode field.
func (o *Event) SetSmtpCode(v int32) {
	o.SmtpCode = &v
}

// GetSmtpDescription returns the SmtpDescription field value if set, zero value otherwise.
func (o *Event) GetSmtpDescription() string {
	if o == nil || IsNil(o.SmtpDescription) {
		var ret string
		return ret
	}
	return *o.SmtpDescription
}

// GetSmtpDescriptionOk returns a tuple with the SmtpDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSmtpDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpDescription) {
		return nil, false
	}
	return o.SmtpDescription, true
}

// HasSmtpDescription returns a boolean if a field has been set.
func (o *Event) HasSmtpDescription() bool {
	if o != nil && !IsNil(o.SmtpDescription) {
		return true
	}

	return false
}

// SetSmtpDescription gets a reference to the given string and assigns it to the SmtpDescription field.
func (o *Event) SetSmtpDescription(v string) {
	o.SmtpDescription = &v
}

// GetEventMetadata returns the EventMetadata field value if set, zero value otherwise.
func (o *Event) GetEventMetadata() EventMetadata {
	if o == nil || IsNil(o.EventMetadata) {
		var ret EventMetadata
		return ret
	}
	return *o.EventMetadata
}

// GetEventMetadataOk returns a tuple with the EventMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventMetadataOk() (*EventMetadata, bool) {
	if o == nil || IsNil(o.EventMetadata) {
		return nil, false
	}
	return o.EventMetadata, true
}

// HasEventMetadata returns a boolean if a field has been set.
func (o *Event) HasEventMetadata() bool {
	if o != nil && !IsNil(o.EventMetadata) {
		return true
	}

	return false
}

// SetEventMetadata gets a reference to the given EventMetadata and assigns it to the EventMetadata field.
func (o *Event) SetEventMetadata(v EventMetadata) {
	o.EventMetadata = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventID) {
		toSerialize["eventID"] = o.EventID
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.IpID) {
		toSerialize["ipID"] = o.IpID
	}
	if !IsNil(o.IpPoolID) {
		toSerialize["ipPoolID"] = o.IpPoolID
	}
	if !IsNil(o.DomainID) {
		toSerialize["domainID"] = o.DomainID
	}
	if !IsNil(o.TpspId) {
		toSerialize["tpspId"] = o.TpspId
	}
	if !IsNil(o.MessageType) {
		toSerialize["messageType"] = o.MessageType
	}
	if !IsNil(o.MessageSubject) {
		toSerialize["messageSubject"] = o.MessageSubject
	}
	if !IsNil(o.AccountID) {
		toSerialize["accountID"] = o.AccountID
	}
	if !IsNil(o.SubAccountID) {
		toSerialize["subAccountID"] = o.SubAccountID
	}
	if !IsNil(o.MessageID) {
		toSerialize["messageID"] = o.MessageID
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.FromName) {
		toSerialize["fromName"] = o.FromName
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.ToName) {
		toSerialize["toName"] = o.ToName
	}
	if !IsNil(o.SubmittedAt) {
		toSerialize["submittedAt"] = o.SubmittedAt
	}
	if !IsNil(o.SmtpCode) {
		toSerialize["smtpCode"] = o.SmtpCode
	}
	if !IsNil(o.SmtpDescription) {
		toSerialize["smtpDescription"] = o.SmtpDescription
	}
	if !IsNil(o.EventMetadata) {
		toSerialize["eventMetadata"] = o.EventMetadata
	}
	return toSerialize, nil
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


