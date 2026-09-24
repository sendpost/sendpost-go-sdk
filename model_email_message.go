/*
SendPost API

# Introduction  > ### 📌 API versioning & the v1 response contract > > This reference documents the **v1 response contract** — the stable, camelCase > response shape that SendPost commits to. This is the shape you should build against. > > **During the current deprecation window**, requests authenticated with an account > or sub-account API key receive the **legacy** response shape by default, so existing > integrations keep working unchanged. To receive the documented v1 shape today, send: > > ``` > X-SendPost-Public-Contract: v1 > ``` > > **How to tell which shape you got.** Every public response echoes the applied > contract in the `X-SendPost-Public-Contract` response header. While the legacy > shape is being served, responses also carry standard deprecation signals: > `Deprecation: true`, a `Sunset` header with the exact cut-over date, and a > `Link: <...>; rel=\"deprecation\"` header pointing at the migration guide. **Read the > `Sunset` header for the authoritative end date** rather than hardcoding one. > > **After the sunset date**, v1 becomes the default and the legacy shape is no longer > served. New integrations should send `X-SendPost-Public-Contract: v1` now and rely on > the shapes in this reference.  SendPost provides email API and SMTP relay which can be used not just to send & measure but also alert & optimised email sending.  You can use SendPost to:  * Send personalised emails to multiple recipients using email API   * Track opens and clicks  * Analyse statistics around open, clicks, bounce, unsubscribe and spam    At and advanced level you can use it to:  * Manage multiple sub-accounts which may map to your promotional or transactional sending, multiple product lines or multiple customers   * Classify your emails using groups for better analysis  * Analyse and fix email sending at sub-account level, IP Pool level or group level  * Have automated alerts to notify disruptions regarding email sending  * Manage different dedicated IP Pools so to better control your email sending  * Automatically know when IP or domain is blacklisted or sender score is down  * Leverage pro deliverability tools to get significantly better email deliverability & inboxing   [<img src=\"https://run.pstmn.io/button.svg\" alt=\"Run In Postman\" style=\"width: 128px; height: 32px;\">](https://god.gw.postman.com/run-collection/33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b%26entityType%3Dcollection%26workspaceId%3D6b1e4f65-96a9-4136-9512-6266c852517e)   # Overview  ## REST API  SendPost API is built on REST API principles. Authenticated users can interact with any of the API endpoints to perform:  * **GET**- to get a resource  * **POST** - to create a resource  * **PUT** - to update an existing resource  * **DELETE** - to delete a resource   The API endpoint for all API calls is: <code>https://api.sendpost.io/api/v1</code>   Some conventions that have been followed in the API design overall are following:   * All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header.  * All resource endpoints end with singular name and not plural. So we have <code>domain</code> instead of domains for domain resource endpoint. Likewise we have <code>sender</code> instead of senders for sender resource endpoint.  * Body submitted for POST / PUT API calls as well as JSON response from SendPost API follow camelcase convention  * All timestamps returned in response (created or submittedAt response fields) are UNIX nano epoch timestamp.   <aside class=\"success\"> All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header. </aside>   SendPost uses conventional HTTP response codes to indicate the success or failure of an API request.    * Codes in the <code>2xx</code> range indicate success.   * Codes in the <code>4xx</code> range indicate an error owing due to unauthorize access, incorrect request parameters or body etc.  * Code in the <code>5xx</code> range indicate an eror with SendPost's servers ( internal service issue or maintenance )   <aside class=\"info\"> SendPost all responses return <code>created</code> in UNIX nano epoch timestamp.  </aside>   ## Authentication  SendPost uses API keys for authentication. You can register a new SendPost API key at our [developer portal](https://app.sendpost.io/register).   SendPost expects the API key to be included in all API requests to the server in a header that looks like the following:   `X-SubAccount-ApiKey: AHEZEP8192SEGH`   This API key is used for all Sub-Account level operations such as:  * Sending emails  * Retrieving stats regarding open, click, bounce, unsubscribe and spam  * Uploading suppressions list  * Verifying sending domains and more  In addition to <code>X-SubAccount-ApiKey</code> you also have another API Key <code>X-Account-APIKey</code> which is used for Account level operations such as :  * Creating and managing sub-accounts  * Allocating IPs for your account  * Getting overall billing and usage information  * Email List validation  * Creating and managing alerts and more   <aside class=\"notice\"> You must look at individual API reference page to look at whether <code>X-SubAccount-ApiKey</code> is required or <code>X-Account-ApiKey</code> </aside>   In case an incorrect API Key header is specified or if it is missed you will get HTTP Response 401 ( Unauthorized ) response from SendPost.   ## HTTP Response Headers   Code           | Reason                 | Details ---------------| -----------------------| ----------- 200            | Success                | Everything went well 401            | Unauthorized           | Incorrect or missing API header either <code>X-SubAccount-ApiKey</code> or <code>X-Account-ApiKey</code> 403            | Forbidden              | Typically sent when resource with same name or details already exist 406            | Missing resource id    | Resource id specified is either missing or doesn't exist 422            | Unprocessable entity   | Request body is not in proper format 500            | Internal server error  | Some error happened at SendPost while processing API request 503            | Service Unavailable    | SendPost is offline for maintenance. Please try again later  # API SDKs  We have native SendPost SDKs in the following programming languages. You can integrate with them or create your own SDK with our API specification. In case you need any assistance with respect to API then do reachout to our team from website chat or email us at **hello@sendpost.io**   * [PHP](https://github.com/sendpost/sendpost_php_sdk)  * [Javascript](https://github.com/sendpost/sendpost_javascript_sdk)  * [Ruby](https://github.com/sendpost/sendpost_ruby_sdk)  * [Python](https://github.com/sendpost/sendpost_python_sdk)  * [Golang](https://github.com/sendpost/sendpost_go_sdk)   # API Reference  SendX REST API can be broken down into two major sub-sections:   * Sub-Account  * Account    Sub-Account API operations enable common email sending API use-cases like sending bulk email, adding new domains or senders for email sending programmatically, retrieving stats, adding suppressions etc. All Sub-Account API operations need to pass <code>X-SubAccount-ApiKey</code> header with every API call.   The Account API operations allow users to manage multiple sub-accounts and manage IPs. A single parent SendPost account can have 100's of sub-accounts. You may want to create sub-accounts for different products your company is running or to segregate types of emails or for managing email sending across multiple customers of yours.   # SMTP Reference  Simple Mail Transfer Protocol (SMTP) is a quick and easy way to send email from one server to another. SendPost provides an SMTP service that allows you to deliver your email via our servers instead of your own client or server.  This means you can count on SendPost's delivery at scale for your SMTP needs.    ## Integrating SMTP    1. Get the SMTP `username` and `password` from your SendPost account.  2. Set the server host in your email client or application to `smtp.sendpost.io`. This setting is sometimes referred to as the external SMTP server or the SMTP relay.  3. Set the `username` and `password`.  4. Set the port to `587` (or as specified below).  ## SMTP Ports   - For an unencrypted or a TLS connection, use port `25`, `2525` or `587`.  - For a SSL connection, use port `465`  - Check your firewall and network to ensure they're not blocking any of our SMTP Endpoints.   SendPost supports STARTTLS for establishing a TLS-encrypted connection. STARTTLS is a means of upgrading an unencrypted connection to an encrypted connection. There are versions of STARTTLS for a variety of protocols; the SMTP version is defined in [RFC 3207](https://www.ietf.org/rfc/rfc3207.txt).   To set up a STARTTLS connection, the SMTP client connects to the SendPost SMTP endpoint `smtp.sendpost.io` on port 25, 587, or 2525, issues an EHLO command, and waits for the server to announce that it supports the STARTTLS SMTP extension. The client then issues the STARTTLS command, initiating TLS negotiation. When negotiation is complete, the client issues an EHLO command over the new encrypted connection, and the SMTP session proceeds normally.   <aside class=\"success\"> If you are unsure which port to use, a TLS connection on port 587 is typically recommended. </aside>   ## Sending email from your application   ```javascript \"use strict\";  const nodemailer = require(\"nodemailer\");  async function main() { // create reusable transporter object using the default SMTP transport let transporter = nodemailer.createTransport({ host: \"smtp.sendpost.io\", port: 587, secure: false, // true for 465, false for other ports auth: { user:  \"<username>\" , // generated ethereal user pass: \"<password>\", // generated ethereal password }, requireTLS: true, debug: true, logger: true, });  // send mail with defined transport object try { let info = await transporter.sendMail({ from: 'erlich@piedpiper.com', to: 'gilfoyle@piedpiper.com', subject: 'Test Email Subject', html: '<h1>Hello Geeks!!!</h1>', }); console.log(\"Message sent: %s\", info.messageId); } catch (e) { console.log(e) } }  main().catch(console.error); ```  For PHP   ```php <?php // Import PHPMailer classes into the global namespace use PHPMailer\\PHPMailer\\PHPMailer; use PHPMailer\\PHPMailer\\SMTP; use PHPMailer\\PHPMailer\\Exception;  // Load Composer's autoloader require 'vendor/autoload.php';  $mail = new PHPMailer(true);  // Settings try { $mail->SMTPDebug = SMTP::DEBUG_CONNECTION;                  // Enable verbose debug output $mail->isSMTP();                                            // Send using SMTP $mail->Host       = 'smtp.sendpost.io';                     // Set the SMTP server to send through $mail->SMTPAuth   = true;                                   // Enable SMTP authentication $mail->Username   = '<username>';                           // SMTP username $mail->Password   = '<password>';                           // SMTP password $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;         // Enable implicit TLS encryption $mail->Port       = 587;                                    // TCP port to connect to; use 587 if you have set `SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS`  //Recipients $mail->setFrom('erlich@piedpiper.com', 'Erlich'); $mail->addAddress('gilfoyle@piedpiper.com', 'Gilfoyle');  //Content $mail->isHTML(true);                                  //Set email format to HTML $mail->Subject = 'Here is the subject'; $mail->Body    = 'This is the HTML message body <b>in bold!</b>'; $mail->AltBody = 'This is the body in plain text for non-HTML mail clients';  $mail->send(); echo 'Message has been sent';  } catch (Exception $e) { echo \"Message could not be sent. Mailer Error: {$mail->ErrorInfo}\"; } ``` For Python ```python #!/usr/bin/python3  import sys import os import re  from smtplib import SMTP import ssl  from email.mime.text import MIMEText  SMTPserver = 'smtp.sendpost.io' PORT = 587 sender =     'erlich@piedpiper.com' destination = ['gilfoyle@piedpiper.com']  USERNAME = \"<username>\" PASSWORD = \"<password>\"  # typical values for text_subtype are plain, html, xml text_subtype = 'plain'  content=\"\"\"\\ Test message \"\"\"  subject=\"Sent from Python\"  try: msg = MIMEText(content, text_subtype) msg['Subject']= subject msg['From']   = sender  conn = SMTP(SMTPserver, PORT) conn.ehlo() context = ssl.create_default_context() conn.starttls(context=context)  # upgrade to tls conn.ehlo() conn.set_debuglevel(True) conn.login(USERNAME, PASSWORD)  try: resp = conn.sendmail(sender, destination, msg.as_string()) print(\"Send Mail Response: \", resp) except Exception as e: print(\"Send Email Error: \", e) finally: conn.quit()  except Exception as e: print(\"Error:\", e) ``` For Golang ```go package main  import ( \"fmt\" \"net/smtp\" \"os\" )  // Sending Email Using Smtp in Golang  func main() {  username := \"<username>\" password := \"<password>\"  from := \"erlich@piedpiper.com\" toList := []string{\"gilfoyle@piedpiper.com\"} host := \"smtp.sendpost.io\" port := \"587\" // recommended  // This is the message to send in the mail msg := \"Hello geeks!!!\"  // We can't send strings directly in mail, // strings need to be converted into slice bytes body := []byte(msg)  // PlainAuth uses the given username and password to // authenticate to host and act as identity. // Usually identity should be the empty string, // to act as username. auth := smtp.PlainAuth(\"\", username, password, host)  // SendMail uses TLS connection to send the mail // The email is sent to all address in the toList, // the body should be of type bytes, not strings // This returns error if any occured. err := smtp.SendMail(host+\":\"+port, auth, from, toList, body)  // handling the errors if err != nil { fmt.Println(err) os.Exit(1) }  fmt.Println(\"Successfully sent mail to all user in toList\") }  ``` For Java ```java // implementation 'com.sun.mail:javax.mail:1.6.2'  import java.util.Properties;  import javax.mail.Message; import javax.mail.Session; import javax.mail.Transport; import javax.mail.internet.InternetAddress; import javax.mail.internet.MimeMessage;  public class SMTPConnect {  // This address must be verified. static final String FROM = \"erlich@piedpiper.com\"; static final String FROMNAME = \"Erlich Bachman\";  // Replace recipient@example.com with a \"To\" address. If your account // is still in the sandbox, this address must be verified. static final String TO = \"gilfoyle@piedpiper.com\";  // Replace smtp_username with your SendPost SMTP user name. static final String SMTP_USERNAME = \"<username>\";  // Replace smtp_password with your SendPost SMTP password. static final String SMTP_PASSWORD = \"<password>\";  // SMTP Host Name static final String HOST = \"smtp.sendpost.io\";  // The port you will connect to on SendPost SMTP Endpoint. static final int PORT = 587;  static final String SUBJECT = \"SendPost SMTP Test (SMTP interface accessed using Java)\";  static final String BODY = String.join( System.getProperty(\"line.separator\"), \"<h1>SendPost SMTP Test</h1>\", \"<p>This email was sent with SendPost using the \", \"<a href='https://github.com/eclipse-ee4j/mail'>Javamail Package</a>\", \" for <a href='https://www.java.com'>Java</a>.\" );  public static void main(String[] args) throws Exception {  // Create a Properties object to contain connection configuration information. Properties props = System.getProperties(); props.put(\"mail.transport.protocol\", \"smtp\"); props.put(\"mail.smtp.port\", PORT); props.put(\"mail.smtp.starttls.enable\", \"true\"); props.put(\"mail.smtp.debug\", \"true\"); props.put(\"mail.smtp.auth\", \"true\");  // Create a Session object to represent a mail session with the specified properties. Session session = Session.getDefaultInstance(props);  // Create a message with the specified information. MimeMessage msg = new MimeMessage(session); msg.setFrom(new InternetAddress(FROM,FROMNAME)); msg.setRecipient(Message.RecipientType.TO, new InternetAddress(TO)); msg.setSubject(SUBJECT); msg.setContent(BODY,\"text/html\");  // Create a transport. Transport transport = session.getTransport();  // Send the message. try { System.out.println(\"Sending...\");  // Connect to SendPost SMTP using the SMTP username and password you specified above. transport.connect(HOST, SMTP_USERNAME, SMTP_PASSWORD);  // Send the email. transport.sendMessage(msg, msg.getAllRecipients()); System.out.println(\"Email sent!\");  } catch (Exception ex) {  System.out.println(\"The email was not sent.\"); System.out.println(\"Error message: \" + ex.getMessage()); System.out.println(ex); } // Close and terminate the connection. } } ```  Many programming languages support sending email using SMTP. This capability might be built into the programming language itself, or it might be available as an add-on, plug-in, or library. You can take advantage of this capability by sending email through SendPost from within application programs that you write.  We have provided examples in Python3, Golang, Java, PHP, JS.  # API Contract Versioning (Public REST)  The public REST API uses a versioned response contract so field changes stay non-breaking:  * Send `X-SendPost-Public-Contract: v1` to opt into the current v1 response shape, or `legacy` for the pre-v1 shape. If the header is omitted, the applied contract is policy-driven — `legacy` before the published sunset date, `v1` after it. * Every response echoes `X-SendPost-Public-Contract: <applied>`. When the `legacy` contract is served, responses also include `Deprecation: true`, `Sunset: <RFC1123 date>`, and `Link: <doc-url>; rel=\"deprecation\"`. * Migrate to `v1` before the sunset date. Notable legacy → v1 field changes: Suppression `smtp_error` → `smtpError`, Stat `email_type` → `emailType`.  > `X-SendPost-Private-Api: true` is an internal header used only by the SendPost dashboard to receive richer internal objects. It is not part of the public SDK contract and should not be set by API integrations. 

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
)

// checks if the EmailMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailMessage{}

// EmailMessage Complete email message details included in webhook payloads. This contains all the original email data plus routing information. 
type EmailMessage struct {
	// Unique identifier for this email message
	MessageID *string `json:"messageID,omitempty"`
	// ID of the SendPost account that sent this email
	AccountID *int64 `json:"accountID,omitempty"`
	// ID of the sub-account that sent this email
	SubAccountID *int64 `json:"subAccountID,omitempty"`
	// ID of the dedicated IP used for sending (0 if a shared IP was used)
	IpID *int64 `json:"ipID,omitempty"`
	// The public IP address used to send this email
	PublicIP *string `json:"publicIP,omitempty"`
	// The internal/local IP address used to send this email
	LocalIP *string `json:"localIP,omitempty"`
	// Classification of the email based on recipient domain: - `gmail` - Gmail recipient - `yahoo` - Yahoo recipient - `microsoft` - Outlook/Hotmail recipient - `default` - Other email providers 
	EmailType *string `json:"emailType,omitempty"`
	// UNIX epoch timestamp in nanoseconds when the email was submitted
	SubmittedAt *int64 `json:"submittedAt,omitempty"`
	// The sender's email address
	From *EmailAddress `json:"from,omitempty"`
	// The reply-to address (if different from sender)
	ReplyTo *EmailAddress `json:"replyTo,omitempty"`
	// The envelope recipient (actual delivery address)
	To *Recipient `json:"to,omitempty"`
	// Tags/groups for categorization and analytics
	Groups []string `json:"groups,omitempty"`
	// Name of the IP pool used for sending
	IpPool *string `json:"ipPool,omitempty"`
	// Custom headers included in the email
	Headers map[string]string `json:"headers,omitempty"`
	// Custom fields sent with the email, available for personalization
	CustomFields map[string]interface{} `json:"customFields,omitempty"`
	// Whether open tracking was enabled
	TrackOpens *bool `json:"trackOpens,omitempty"`
	// Whether click tracking was enabled
	TrackClicks *bool `json:"trackClicks,omitempty"`
	// Custom webhook endpoint for this email (if specified)
	WebhookEndpoint *string `json:"webhookEndpoint,omitempty"`
}

// NewEmailMessage instantiates a new EmailMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailMessage() *EmailMessage {
	this := EmailMessage{}
	return &this
}

// NewEmailMessageWithDefaults instantiates a new EmailMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailMessageWithDefaults() *EmailMessage {
	this := EmailMessage{}
	return &this
}

// GetMessageID returns the MessageID field value if set, zero value otherwise.
func (o *EmailMessage) GetMessageID() string {
	if o == nil || IsNil(o.MessageID) {
		var ret string
		return ret
	}
	return *o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetMessageIDOk() (*string, bool) {
	if o == nil || IsNil(o.MessageID) {
		return nil, false
	}
	return o.MessageID, true
}

// HasMessageID returns a boolean if a field has been set.
func (o *EmailMessage) HasMessageID() bool {
	if o != nil && !IsNil(o.MessageID) {
		return true
	}

	return false
}

// SetMessageID gets a reference to the given string and assigns it to the MessageID field.
func (o *EmailMessage) SetMessageID(v string) {
	o.MessageID = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *EmailMessage) GetAccountID() int64 {
	if o == nil || IsNil(o.AccountID) {
		var ret int64
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetAccountIDOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountID returns a boolean if a field has been set.
func (o *EmailMessage) HasAccountID() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given int64 and assigns it to the AccountID field.
func (o *EmailMessage) SetAccountID(v int64) {
	o.AccountID = &v
}

// GetSubAccountID returns the SubAccountID field value if set, zero value otherwise.
func (o *EmailMessage) GetSubAccountID() int64 {
	if o == nil || IsNil(o.SubAccountID) {
		var ret int64
		return ret
	}
	return *o.SubAccountID
}

// GetSubAccountIDOk returns a tuple with the SubAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetSubAccountIDOk() (*int64, bool) {
	if o == nil || IsNil(o.SubAccountID) {
		return nil, false
	}
	return o.SubAccountID, true
}

// HasSubAccountID returns a boolean if a field has been set.
func (o *EmailMessage) HasSubAccountID() bool {
	if o != nil && !IsNil(o.SubAccountID) {
		return true
	}

	return false
}

// SetSubAccountID gets a reference to the given int64 and assigns it to the SubAccountID field.
func (o *EmailMessage) SetSubAccountID(v int64) {
	o.SubAccountID = &v
}

// GetIpID returns the IpID field value if set, zero value otherwise.
func (o *EmailMessage) GetIpID() int64 {
	if o == nil || IsNil(o.IpID) {
		var ret int64
		return ret
	}
	return *o.IpID
}

// GetIpIDOk returns a tuple with the IpID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetIpIDOk() (*int64, bool) {
	if o == nil || IsNil(o.IpID) {
		return nil, false
	}
	return o.IpID, true
}

// HasIpID returns a boolean if a field has been set.
func (o *EmailMessage) HasIpID() bool {
	if o != nil && !IsNil(o.IpID) {
		return true
	}

	return false
}

// SetIpID gets a reference to the given int64 and assigns it to the IpID field.
func (o *EmailMessage) SetIpID(v int64) {
	o.IpID = &v
}

// GetPublicIP returns the PublicIP field value if set, zero value otherwise.
func (o *EmailMessage) GetPublicIP() string {
	if o == nil || IsNil(o.PublicIP) {
		var ret string
		return ret
	}
	return *o.PublicIP
}

// GetPublicIPOk returns a tuple with the PublicIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetPublicIPOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIP) {
		return nil, false
	}
	return o.PublicIP, true
}

// HasPublicIP returns a boolean if a field has been set.
func (o *EmailMessage) HasPublicIP() bool {
	if o != nil && !IsNil(o.PublicIP) {
		return true
	}

	return false
}

// SetPublicIP gets a reference to the given string and assigns it to the PublicIP field.
func (o *EmailMessage) SetPublicIP(v string) {
	o.PublicIP = &v
}

// GetLocalIP returns the LocalIP field value if set, zero value otherwise.
func (o *EmailMessage) GetLocalIP() string {
	if o == nil || IsNil(o.LocalIP) {
		var ret string
		return ret
	}
	return *o.LocalIP
}

// GetLocalIPOk returns a tuple with the LocalIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetLocalIPOk() (*string, bool) {
	if o == nil || IsNil(o.LocalIP) {
		return nil, false
	}
	return o.LocalIP, true
}

// HasLocalIP returns a boolean if a field has been set.
func (o *EmailMessage) HasLocalIP() bool {
	if o != nil && !IsNil(o.LocalIP) {
		return true
	}

	return false
}

// SetLocalIP gets a reference to the given string and assigns it to the LocalIP field.
func (o *EmailMessage) SetLocalIP(v string) {
	o.LocalIP = &v
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *EmailMessage) GetEmailType() string {
	if o == nil || IsNil(o.EmailType) {
		var ret string
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetEmailTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HasEmailType returns a boolean if a field has been set.
func (o *EmailMessage) HasEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given string and assigns it to the EmailType field.
func (o *EmailMessage) SetEmailType(v string) {
	o.EmailType = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *EmailMessage) GetSubmittedAt() int64 {
	if o == nil || IsNil(o.SubmittedAt) {
		var ret int64
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetSubmittedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.SubmittedAt) {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *EmailMessage) HasSubmittedAt() bool {
	if o != nil && !IsNil(o.SubmittedAt) {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given int64 and assigns it to the SubmittedAt field.
func (o *EmailMessage) SetSubmittedAt(v int64) {
	o.SubmittedAt = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EmailMessage) GetFrom() EmailAddress {
	if o == nil || IsNil(o.From) {
		var ret EmailAddress
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetFromOk() (*EmailAddress, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EmailMessage) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given EmailAddress and assigns it to the From field.
func (o *EmailMessage) SetFrom(v EmailAddress) {
	o.From = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *EmailMessage) GetReplyTo() EmailAddress {
	if o == nil || IsNil(o.ReplyTo) {
		var ret EmailAddress
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetReplyToOk() (*EmailAddress, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *EmailMessage) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given EmailAddress and assigns it to the ReplyTo field.
func (o *EmailMessage) SetReplyTo(v EmailAddress) {
	o.ReplyTo = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EmailMessage) GetTo() Recipient {
	if o == nil || IsNil(o.To) {
		var ret Recipient
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetToOk() (*Recipient, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EmailMessage) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given Recipient and assigns it to the To field.
func (o *EmailMessage) SetTo(v Recipient) {
	o.To = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *EmailMessage) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *EmailMessage) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *EmailMessage) SetGroups(v []string) {
	o.Groups = v
}

// GetIpPool returns the IpPool field value if set, zero value otherwise.
func (o *EmailMessage) GetIpPool() string {
	if o == nil || IsNil(o.IpPool) {
		var ret string
		return ret
	}
	return *o.IpPool
}

// GetIpPoolOk returns a tuple with the IpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetIpPoolOk() (*string, bool) {
	if o == nil || IsNil(o.IpPool) {
		return nil, false
	}
	return o.IpPool, true
}

// HasIpPool returns a boolean if a field has been set.
func (o *EmailMessage) HasIpPool() bool {
	if o != nil && !IsNil(o.IpPool) {
		return true
	}

	return false
}

// SetIpPool gets a reference to the given string and assigns it to the IpPool field.
func (o *EmailMessage) SetIpPool(v string) {
	o.IpPool = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *EmailMessage) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *EmailMessage) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *EmailMessage) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *EmailMessage) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EmailMessage) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *EmailMessage) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTrackOpens returns the TrackOpens field value if set, zero value otherwise.
func (o *EmailMessage) GetTrackOpens() bool {
	if o == nil || IsNil(o.TrackOpens) {
		var ret bool
		return ret
	}
	return *o.TrackOpens
}

// GetTrackOpensOk returns a tuple with the TrackOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetTrackOpensOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackOpens) {
		return nil, false
	}
	return o.TrackOpens, true
}

// HasTrackOpens returns a boolean if a field has been set.
func (o *EmailMessage) HasTrackOpens() bool {
	if o != nil && !IsNil(o.TrackOpens) {
		return true
	}

	return false
}

// SetTrackOpens gets a reference to the given bool and assigns it to the TrackOpens field.
func (o *EmailMessage) SetTrackOpens(v bool) {
	o.TrackOpens = &v
}

// GetTrackClicks returns the TrackClicks field value if set, zero value otherwise.
func (o *EmailMessage) GetTrackClicks() bool {
	if o == nil || IsNil(o.TrackClicks) {
		var ret bool
		return ret
	}
	return *o.TrackClicks
}

// GetTrackClicksOk returns a tuple with the TrackClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetTrackClicksOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackClicks) {
		return nil, false
	}
	return o.TrackClicks, true
}

// HasTrackClicks returns a boolean if a field has been set.
func (o *EmailMessage) HasTrackClicks() bool {
	if o != nil && !IsNil(o.TrackClicks) {
		return true
	}

	return false
}

// SetTrackClicks gets a reference to the given bool and assigns it to the TrackClicks field.
func (o *EmailMessage) SetTrackClicks(v bool) {
	o.TrackClicks = &v
}

// GetWebhookEndpoint returns the WebhookEndpoint field value if set, zero value otherwise.
func (o *EmailMessage) GetWebhookEndpoint() string {
	if o == nil || IsNil(o.WebhookEndpoint) {
		var ret string
		return ret
	}
	return *o.WebhookEndpoint
}

// GetWebhookEndpointOk returns a tuple with the WebhookEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessage) GetWebhookEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEndpoint) {
		return nil, false
	}
	return o.WebhookEndpoint, true
}

// HasWebhookEndpoint returns a boolean if a field has been set.
func (o *EmailMessage) HasWebhookEndpoint() bool {
	if o != nil && !IsNil(o.WebhookEndpoint) {
		return true
	}

	return false
}

// SetWebhookEndpoint gets a reference to the given string and assigns it to the WebhookEndpoint field.
func (o *EmailMessage) SetWebhookEndpoint(v string) {
	o.WebhookEndpoint = &v
}

func (o EmailMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageID) {
		toSerialize["messageID"] = o.MessageID
	}
	if !IsNil(o.AccountID) {
		toSerialize["accountID"] = o.AccountID
	}
	if !IsNil(o.SubAccountID) {
		toSerialize["subAccountID"] = o.SubAccountID
	}
	if !IsNil(o.IpID) {
		toSerialize["ipID"] = o.IpID
	}
	if !IsNil(o.PublicIP) {
		toSerialize["publicIP"] = o.PublicIP
	}
	if !IsNil(o.LocalIP) {
		toSerialize["localIP"] = o.LocalIP
	}
	if !IsNil(o.EmailType) {
		toSerialize["emailType"] = o.EmailType
	}
	if !IsNil(o.SubmittedAt) {
		toSerialize["submittedAt"] = o.SubmittedAt
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.IpPool) {
		toSerialize["ipPool"] = o.IpPool
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.TrackOpens) {
		toSerialize["trackOpens"] = o.TrackOpens
	}
	if !IsNil(o.TrackClicks) {
		toSerialize["trackClicks"] = o.TrackClicks
	}
	if !IsNil(o.WebhookEndpoint) {
		toSerialize["webhookEndpoint"] = o.WebhookEndpoint
	}
	return toSerialize, nil
}

type NullableEmailMessage struct {
	value *EmailMessage
	isSet bool
}

func (v NullableEmailMessage) Get() *EmailMessage {
	return v.value
}

func (v *NullableEmailMessage) Set(val *EmailMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailMessage(val *EmailMessage) *NullableEmailMessage {
	return &NullableEmailMessage{value: val, isSet: true}
}

func (v NullableEmailMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


