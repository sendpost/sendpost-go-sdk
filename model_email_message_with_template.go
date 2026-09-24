/*
SendPost API

# Introduction  > ### 📌 API versioning & the v1 response contract > > This reference documents the **v1 response contract** — the stable, camelCase > response shape that SendPost commits to. This is the shape you should build against. > > **During the current deprecation window**, requests authenticated with an account > or sub-account API key receive the **legacy** response shape by default, so existing > integrations keep working unchanged. To receive the documented v1 shape today, send: > > ``` > X-SendPost-Public-Contract: v1 > ``` > > **How to tell which shape you got.** Every public response echoes the applied > contract in the `X-SendPost-Public-Contract` response header. While the legacy > shape is being served, responses also carry standard deprecation signals: > `Deprecation: true`, a `Sunset` header with the exact cut-over date, and a > `Link: <...>; rel=\"deprecation\"` header pointing at the migration guide. **Read the > `Sunset` header for the authoritative end date** rather than hardcoding one. > > **After the sunset date**, v1 becomes the default and the legacy shape is no longer > served. New integrations should send `X-SendPost-Public-Contract: v1` now and rely on > the shapes in this reference.  SendPost provides email API and SMTP relay which can be used not just to send & measure but also alert & optimised email sending.  You can use SendPost to:  * Send personalised emails to multiple recipients using email API   * Track opens and clicks  * Analyse statistics around open, clicks, bounce, unsubscribe and spam    At and advanced level you can use it to:  * Manage multiple sub-accounts which may map to your promotional or transactional sending, multiple product lines or multiple customers   * Classify your emails using groups for better analysis  * Analyse and fix email sending at sub-account level, IP Pool level or group level  * Have automated alerts to notify disruptions regarding email sending  * Manage different dedicated IP Pools so to better control your email sending  * Automatically know when IP or domain is blacklisted or sender score is down  * Leverage pro deliverability tools to get significantly better email deliverability & inboxing   [<img src=\"https://run.pstmn.io/button.svg\" alt=\"Run In Postman\" style=\"width: 128px; height: 32px;\">](https://god.gw.postman.com/run-collection/33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b%26entityType%3Dcollection%26workspaceId%3D6b1e4f65-96a9-4136-9512-6266c852517e)   # Overview  ## REST API  SendPost API is built on REST API principles. Authenticated users can interact with any of the API endpoints to perform:  * **GET**- to get a resource  * **POST** - to create a resource  * **PUT** - to update an existing resource  * **DELETE** - to delete a resource   The API endpoint for all API calls is: <code>https://api.sendpost.io/api/v1</code>   Some conventions that have been followed in the API design overall are following:   * All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header.  * All resource endpoints end with singular name and not plural. So we have <code>domain</code> instead of domains for domain resource endpoint. Likewise we have <code>sender</code> instead of senders for sender resource endpoint.  * Body submitted for POST / PUT API calls as well as JSON response from SendPost API follow camelcase convention  * All timestamps returned in response (created or submittedAt response fields) are UNIX nano epoch timestamp.   <aside class=\"success\"> All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header. </aside>   SendPost uses conventional HTTP response codes to indicate the success or failure of an API request.    * Codes in the <code>2xx</code> range indicate success.   * Codes in the <code>4xx</code> range indicate an error owing due to unauthorize access, incorrect request parameters or body etc.  * Code in the <code>5xx</code> range indicate an eror with SendPost's servers ( internal service issue or maintenance )   <aside class=\"info\"> SendPost all responses return <code>created</code> in UNIX nano epoch timestamp.  </aside>   ## Authentication  SendPost uses API keys for authentication. You can register a new SendPost API key at our [developer portal](https://app.sendpost.io/register).   SendPost expects the API key to be included in all API requests to the server in a header that looks like the following:   `X-SubAccount-ApiKey: AHEZEP8192SEGH`   This API key is used for all Sub-Account level operations such as:  * Sending emails  * Retrieving stats regarding open, click, bounce, unsubscribe and spam  * Uploading suppressions list  * Verifying sending domains and more  In addition to <code>X-SubAccount-ApiKey</code> you also have another API Key <code>X-Account-APIKey</code> which is used for Account level operations such as :  * Creating and managing sub-accounts  * Allocating IPs for your account  * Getting overall billing and usage information  * Email List validation  * Creating and managing alerts and more   <aside class=\"notice\"> You must look at individual API reference page to look at whether <code>X-SubAccount-ApiKey</code> is required or <code>X-Account-ApiKey</code> </aside>   In case an incorrect API Key header is specified or if it is missed you will get HTTP Response 401 ( Unauthorized ) response from SendPost.   ## HTTP Response Headers   Code           | Reason                 | Details ---------------| -----------------------| ----------- 200            | Success                | Everything went well 401            | Unauthorized           | Incorrect or missing API header either <code>X-SubAccount-ApiKey</code> or <code>X-Account-ApiKey</code> 403            | Forbidden              | Typically sent when resource with same name or details already exist 406            | Missing resource id    | Resource id specified is either missing or doesn't exist 422            | Unprocessable entity   | Request body is not in proper format 500            | Internal server error  | Some error happened at SendPost while processing API request 503            | Service Unavailable    | SendPost is offline for maintenance. Please try again later  # API SDKs  We have native SendPost SDKs in the following programming languages. You can integrate with them or create your own SDK with our API specification. In case you need any assistance with respect to API then do reachout to our team from website chat or email us at **hello@sendpost.io**   * [PHP](https://github.com/sendpost/sendpost_php_sdk)  * [Javascript](https://github.com/sendpost/sendpost_javascript_sdk)  * [Ruby](https://github.com/sendpost/sendpost_ruby_sdk)  * [Python](https://github.com/sendpost/sendpost_python_sdk)  * [Golang](https://github.com/sendpost/sendpost_go_sdk)   # API Reference  SendX REST API can be broken down into two major sub-sections:   * Sub-Account  * Account    Sub-Account API operations enable common email sending API use-cases like sending bulk email, adding new domains or senders for email sending programmatically, retrieving stats, adding suppressions etc. All Sub-Account API operations need to pass <code>X-SubAccount-ApiKey</code> header with every API call.   The Account API operations allow users to manage multiple sub-accounts and manage IPs. A single parent SendPost account can have 100's of sub-accounts. You may want to create sub-accounts for different products your company is running or to segregate types of emails or for managing email sending across multiple customers of yours.   # SMTP Reference  Simple Mail Transfer Protocol (SMTP) is a quick and easy way to send email from one server to another. SendPost provides an SMTP service that allows you to deliver your email via our servers instead of your own client or server.  This means you can count on SendPost's delivery at scale for your SMTP needs.    ## Integrating SMTP    1. Get the SMTP `username` and `password` from your SendPost account.  2. Set the server host in your email client or application to `smtp.sendpost.io`. This setting is sometimes referred to as the external SMTP server or the SMTP relay.  3. Set the `username` and `password`.  4. Set the port to `587` (or as specified below).  ## SMTP Ports   - For an unencrypted or a TLS connection, use port `25`, `2525` or `587`.  - For a SSL connection, use port `465`  - Check your firewall and network to ensure they're not blocking any of our SMTP Endpoints.   SendPost supports STARTTLS for establishing a TLS-encrypted connection. STARTTLS is a means of upgrading an unencrypted connection to an encrypted connection. There are versions of STARTTLS for a variety of protocols; the SMTP version is defined in [RFC 3207](https://www.ietf.org/rfc/rfc3207.txt).   To set up a STARTTLS connection, the SMTP client connects to the SendPost SMTP endpoint `smtp.sendpost.io` on port 25, 587, or 2525, issues an EHLO command, and waits for the server to announce that it supports the STARTTLS SMTP extension. The client then issues the STARTTLS command, initiating TLS negotiation. When negotiation is complete, the client issues an EHLO command over the new encrypted connection, and the SMTP session proceeds normally.   <aside class=\"success\"> If you are unsure which port to use, a TLS connection on port 587 is typically recommended. </aside>   ## Sending email from your application   ```javascript \"use strict\";  const nodemailer = require(\"nodemailer\");  async function main() { // create reusable transporter object using the default SMTP transport let transporter = nodemailer.createTransport({ host: \"smtp.sendpost.io\", port: 587, secure: false, // true for 465, false for other ports auth: { user:  \"<username>\" , // generated ethereal user pass: \"<password>\", // generated ethereal password }, requireTLS: true, debug: true, logger: true, });  // send mail with defined transport object try { let info = await transporter.sendMail({ from: 'erlich@piedpiper.com', to: 'gilfoyle@piedpiper.com', subject: 'Test Email Subject', html: '<h1>Hello Geeks!!!</h1>', }); console.log(\"Message sent: %s\", info.messageId); } catch (e) { console.log(e) } }  main().catch(console.error); ```  For PHP   ```php <?php // Import PHPMailer classes into the global namespace use PHPMailer\\PHPMailer\\PHPMailer; use PHPMailer\\PHPMailer\\SMTP; use PHPMailer\\PHPMailer\\Exception;  // Load Composer's autoloader require 'vendor/autoload.php';  $mail = new PHPMailer(true);  // Settings try { $mail->SMTPDebug = SMTP::DEBUG_CONNECTION;                  // Enable verbose debug output $mail->isSMTP();                                            // Send using SMTP $mail->Host       = 'smtp.sendpost.io';                     // Set the SMTP server to send through $mail->SMTPAuth   = true;                                   // Enable SMTP authentication $mail->Username   = '<username>';                           // SMTP username $mail->Password   = '<password>';                           // SMTP password $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;         // Enable implicit TLS encryption $mail->Port       = 587;                                    // TCP port to connect to; use 587 if you have set `SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS`  //Recipients $mail->setFrom('erlich@piedpiper.com', 'Erlich'); $mail->addAddress('gilfoyle@piedpiper.com', 'Gilfoyle');  //Content $mail->isHTML(true);                                  //Set email format to HTML $mail->Subject = 'Here is the subject'; $mail->Body    = 'This is the HTML message body <b>in bold!</b>'; $mail->AltBody = 'This is the body in plain text for non-HTML mail clients';  $mail->send(); echo 'Message has been sent';  } catch (Exception $e) { echo \"Message could not be sent. Mailer Error: {$mail->ErrorInfo}\"; } ``` For Python ```python #!/usr/bin/python3  import sys import os import re  from smtplib import SMTP import ssl  from email.mime.text import MIMEText  SMTPserver = 'smtp.sendpost.io' PORT = 587 sender =     'erlich@piedpiper.com' destination = ['gilfoyle@piedpiper.com']  USERNAME = \"<username>\" PASSWORD = \"<password>\"  # typical values for text_subtype are plain, html, xml text_subtype = 'plain'  content=\"\"\"\\ Test message \"\"\"  subject=\"Sent from Python\"  try: msg = MIMEText(content, text_subtype) msg['Subject']= subject msg['From']   = sender  conn = SMTP(SMTPserver, PORT) conn.ehlo() context = ssl.create_default_context() conn.starttls(context=context)  # upgrade to tls conn.ehlo() conn.set_debuglevel(True) conn.login(USERNAME, PASSWORD)  try: resp = conn.sendmail(sender, destination, msg.as_string()) print(\"Send Mail Response: \", resp) except Exception as e: print(\"Send Email Error: \", e) finally: conn.quit()  except Exception as e: print(\"Error:\", e) ``` For Golang ```go package main  import ( \"fmt\" \"net/smtp\" \"os\" )  // Sending Email Using Smtp in Golang  func main() {  username := \"<username>\" password := \"<password>\"  from := \"erlich@piedpiper.com\" toList := []string{\"gilfoyle@piedpiper.com\"} host := \"smtp.sendpost.io\" port := \"587\" // recommended  // This is the message to send in the mail msg := \"Hello geeks!!!\"  // We can't send strings directly in mail, // strings need to be converted into slice bytes body := []byte(msg)  // PlainAuth uses the given username and password to // authenticate to host and act as identity. // Usually identity should be the empty string, // to act as username. auth := smtp.PlainAuth(\"\", username, password, host)  // SendMail uses TLS connection to send the mail // The email is sent to all address in the toList, // the body should be of type bytes, not strings // This returns error if any occured. err := smtp.SendMail(host+\":\"+port, auth, from, toList, body)  // handling the errors if err != nil { fmt.Println(err) os.Exit(1) }  fmt.Println(\"Successfully sent mail to all user in toList\") }  ``` For Java ```java // implementation 'com.sun.mail:javax.mail:1.6.2'  import java.util.Properties;  import javax.mail.Message; import javax.mail.Session; import javax.mail.Transport; import javax.mail.internet.InternetAddress; import javax.mail.internet.MimeMessage;  public class SMTPConnect {  // This address must be verified. static final String FROM = \"erlich@piedpiper.com\"; static final String FROMNAME = \"Erlich Bachman\";  // Replace recipient@example.com with a \"To\" address. If your account // is still in the sandbox, this address must be verified. static final String TO = \"gilfoyle@piedpiper.com\";  // Replace smtp_username with your SendPost SMTP user name. static final String SMTP_USERNAME = \"<username>\";  // Replace smtp_password with your SendPost SMTP password. static final String SMTP_PASSWORD = \"<password>\";  // SMTP Host Name static final String HOST = \"smtp.sendpost.io\";  // The port you will connect to on SendPost SMTP Endpoint. static final int PORT = 587;  static final String SUBJECT = \"SendPost SMTP Test (SMTP interface accessed using Java)\";  static final String BODY = String.join( System.getProperty(\"line.separator\"), \"<h1>SendPost SMTP Test</h1>\", \"<p>This email was sent with SendPost using the \", \"<a href='https://github.com/eclipse-ee4j/mail'>Javamail Package</a>\", \" for <a href='https://www.java.com'>Java</a>.\" );  public static void main(String[] args) throws Exception {  // Create a Properties object to contain connection configuration information. Properties props = System.getProperties(); props.put(\"mail.transport.protocol\", \"smtp\"); props.put(\"mail.smtp.port\", PORT); props.put(\"mail.smtp.starttls.enable\", \"true\"); props.put(\"mail.smtp.debug\", \"true\"); props.put(\"mail.smtp.auth\", \"true\");  // Create a Session object to represent a mail session with the specified properties. Session session = Session.getDefaultInstance(props);  // Create a message with the specified information. MimeMessage msg = new MimeMessage(session); msg.setFrom(new InternetAddress(FROM,FROMNAME)); msg.setRecipient(Message.RecipientType.TO, new InternetAddress(TO)); msg.setSubject(SUBJECT); msg.setContent(BODY,\"text/html\");  // Create a transport. Transport transport = session.getTransport();  // Send the message. try { System.out.println(\"Sending...\");  // Connect to SendPost SMTP using the SMTP username and password you specified above. transport.connect(HOST, SMTP_USERNAME, SMTP_PASSWORD);  // Send the email. transport.sendMessage(msg, msg.getAllRecipients()); System.out.println(\"Email sent!\");  } catch (Exception ex) {  System.out.println(\"The email was not sent.\"); System.out.println(\"Error message: \" + ex.getMessage()); System.out.println(ex); } // Close and terminate the connection. } } ```  Many programming languages support sending email using SMTP. This capability might be built into the programming language itself, or it might be available as an add-on, plug-in, or library. You can take advantage of this capability by sending email through SendPost from within application programs that you write.  We have provided examples in Python3, Golang, Java, PHP, JS.  # API Contract Versioning (Public REST)  The public REST API uses a versioned response contract so field changes stay non-breaking:  * Send `X-SendPost-Public-Contract: v1` to opt into the current v1 response shape, or `legacy` for the pre-v1 shape. If the header is omitted, the applied contract is policy-driven — `legacy` before the published sunset date, `v1` after it. * Every response echoes `X-SendPost-Public-Contract: <applied>`. When the `legacy` contract is served, responses also include `Deprecation: true`, `Sunset: <RFC1123 date>`, and `Link: <doc-url>; rel=\"deprecation\"`. * Migrate to `v1` before the sunset date. Notable legacy → v1 field changes: Suppression `smtp_error` → `smtpError`, Stat `email_type` → `emailType`.  > `X-SendPost-Private-Api: true` is an internal header used only by the SendPost dashboard to receive richer internal objects. It is not part of the public SDK contract and should not be set by API integrations. 

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EmailMessageWithTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailMessageWithTemplate{}

// EmailMessageWithTemplate struct for EmailMessageWithTemplate
type EmailMessageWithTemplate struct {
	// The sender's email address and optional display name
	From EmailAddress `json:"from"`
	// The reply-to email address. If not specified, replies will go to the `from` address
	ReplyTo *EmailAddress `json:"replyTo,omitempty"`
	// List of recipients. Each recipient can have their own CC, BCC, and custom fields for personalization. Maximum 1000 recipients per API call. 
	To []Recipient `json:"to"`
	// Email subject line. Supports Handlebars templating for personalization. Example: \"Hello, {{firstName}}! Your order is ready\" 
	Subject *string `json:"subject,omitempty"`
	// Preview text (preheader) shown in email clients before opening the email. This text appears after the subject line in most email clients' inbox view. 
	PreText *string `json:"preText,omitempty"`
	// HTML content of the email. Supports Handlebars templating for personalization. Use {{customFieldName}} to insert recipient-specific values. 
	HtmlBody *string `json:"htmlBody,omitempty"`
	// Plain text content of the email. Used as fallback when HTML cannot be rendered. Also improves deliverability as some spam filters prefer multipart emails. 
	TextBody *string `json:"textBody,omitempty"`
	// AMP HTML content for supported email clients (Gmail, Yahoo). Enables interactive email experiences like carousels, forms, and real-time content. See https://amp.dev/about/email/ for more details. 
	AmpBody *string `json:"ampBody,omitempty"`
	Template *string `json:"template,omitempty"`
	// Name of the IP pool to use for sending this email. If not specified, the default IP pool for the sub-account will be used. 
	Ippool *string `json:"ippool,omitempty"`
	// Custom email headers to include in the message. Common uses: adding List-Unsubscribe headers, custom tracking IDs, or priority flags. Note: Some headers like From, To, Subject are set automatically and cannot be overridden. 
	Headers map[string]string `json:"headers,omitempty"`
	// Whether to track email opens using a tracking pixel. When enabled, a 1x1 transparent image is inserted into the HTML body. Default: true (if not specified) 
	TrackOpens *bool `json:"trackOpens,omitempty"`
	// Whether to track link clicks by rewriting URLs through SendPost's tracking domain. When enabled, all links in htmlBody are replaced with tracking URLs. Default: true (if not specified) 
	TrackClicks *bool `json:"trackClicks,omitempty"`
	// Tags/groups to categorize this email for analytics and reporting. Use groups to segment your email statistics (e.g., by campaign, email type, or customer segment). 
	Groups []string `json:"groups,omitempty"`
	// File attachments to include with the email. Maximum total attachment size: 25MB. Supported formats: PDF, images, documents, etc. 
	Attachments []Attachment `json:"attachments,omitempty"`
	// Custom webhook URL to receive events for this specific email. Overrides the default webhook configured at the account level. Useful for per-email or per-customer webhook routing. 
	WebhookEndpoint *string `json:"webhookEndpoint,omitempty"`
	// Template ID for the email template
	TemplateId *string `json:"templateId,omitempty"`
	// Key-Value pair of template variables
	TemplateVariables map[string]string `json:"templateVariables,omitempty"`
}

type _EmailMessageWithTemplate EmailMessageWithTemplate

// NewEmailMessageWithTemplate instantiates a new EmailMessageWithTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailMessageWithTemplate(from EmailAddress, to []Recipient) *EmailMessageWithTemplate {
	this := EmailMessageWithTemplate{}
	this.From = from
	this.To = to
	var trackOpens bool = true
	this.TrackOpens = &trackOpens
	var trackClicks bool = true
	this.TrackClicks = &trackClicks
	return &this
}

// NewEmailMessageWithTemplateWithDefaults instantiates a new EmailMessageWithTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailMessageWithTemplateWithDefaults() *EmailMessageWithTemplate {
	this := EmailMessageWithTemplate{}
	var trackOpens bool = true
	this.TrackOpens = &trackOpens
	var trackClicks bool = true
	this.TrackClicks = &trackClicks
	return &this
}

// GetFrom returns the From field value
func (o *EmailMessageWithTemplate) GetFrom() EmailAddress {
	if o == nil {
		var ret EmailAddress
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetFromOk() (*EmailAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *EmailMessageWithTemplate) SetFrom(v EmailAddress) {
	o.From = v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetReplyTo() EmailAddress {
	if o == nil || IsNil(o.ReplyTo) {
		var ret EmailAddress
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetReplyToOk() (*EmailAddress, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given EmailAddress and assigns it to the ReplyTo field.
func (o *EmailMessageWithTemplate) SetReplyTo(v EmailAddress) {
	o.ReplyTo = &v
}

// GetTo returns the To field value
func (o *EmailMessageWithTemplate) GetTo() []Recipient {
	if o == nil {
		var ret []Recipient
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetToOk() ([]Recipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *EmailMessageWithTemplate) SetTo(v []Recipient) {
	o.To = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailMessageWithTemplate) SetSubject(v string) {
	o.Subject = &v
}

// GetPreText returns the PreText field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetPreText() string {
	if o == nil || IsNil(o.PreText) {
		var ret string
		return ret
	}
	return *o.PreText
}

// GetPreTextOk returns a tuple with the PreText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetPreTextOk() (*string, bool) {
	if o == nil || IsNil(o.PreText) {
		return nil, false
	}
	return o.PreText, true
}

// HasPreText returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasPreText() bool {
	if o != nil && !IsNil(o.PreText) {
		return true
	}

	return false
}

// SetPreText gets a reference to the given string and assigns it to the PreText field.
func (o *EmailMessageWithTemplate) SetPreText(v string) {
	o.PreText = &v
}

// GetHtmlBody returns the HtmlBody field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetHtmlBody() string {
	if o == nil || IsNil(o.HtmlBody) {
		var ret string
		return ret
	}
	return *o.HtmlBody
}

// GetHtmlBodyOk returns a tuple with the HtmlBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetHtmlBodyOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlBody) {
		return nil, false
	}
	return o.HtmlBody, true
}

// HasHtmlBody returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasHtmlBody() bool {
	if o != nil && !IsNil(o.HtmlBody) {
		return true
	}

	return false
}

// SetHtmlBody gets a reference to the given string and assigns it to the HtmlBody field.
func (o *EmailMessageWithTemplate) SetHtmlBody(v string) {
	o.HtmlBody = &v
}

// GetTextBody returns the TextBody field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetTextBody() string {
	if o == nil || IsNil(o.TextBody) {
		var ret string
		return ret
	}
	return *o.TextBody
}

// GetTextBodyOk returns a tuple with the TextBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetTextBodyOk() (*string, bool) {
	if o == nil || IsNil(o.TextBody) {
		return nil, false
	}
	return o.TextBody, true
}

// HasTextBody returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasTextBody() bool {
	if o != nil && !IsNil(o.TextBody) {
		return true
	}

	return false
}

// SetTextBody gets a reference to the given string and assigns it to the TextBody field.
func (o *EmailMessageWithTemplate) SetTextBody(v string) {
	o.TextBody = &v
}

// GetAmpBody returns the AmpBody field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetAmpBody() string {
	if o == nil || IsNil(o.AmpBody) {
		var ret string
		return ret
	}
	return *o.AmpBody
}

// GetAmpBodyOk returns a tuple with the AmpBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetAmpBodyOk() (*string, bool) {
	if o == nil || IsNil(o.AmpBody) {
		return nil, false
	}
	return o.AmpBody, true
}

// HasAmpBody returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasAmpBody() bool {
	if o != nil && !IsNil(o.AmpBody) {
		return true
	}

	return false
}

// SetAmpBody gets a reference to the given string and assigns it to the AmpBody field.
func (o *EmailMessageWithTemplate) SetAmpBody(v string) {
	o.AmpBody = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *EmailMessageWithTemplate) SetTemplate(v string) {
	o.Template = &v
}

// GetIppool returns the Ippool field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetIppool() string {
	if o == nil || IsNil(o.Ippool) {
		var ret string
		return ret
	}
	return *o.Ippool
}

// GetIppoolOk returns a tuple with the Ippool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetIppoolOk() (*string, bool) {
	if o == nil || IsNil(o.Ippool) {
		return nil, false
	}
	return o.Ippool, true
}

// HasIppool returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasIppool() bool {
	if o != nil && !IsNil(o.Ippool) {
		return true
	}

	return false
}

// SetIppool gets a reference to the given string and assigns it to the Ippool field.
func (o *EmailMessageWithTemplate) SetIppool(v string) {
	o.Ippool = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *EmailMessageWithTemplate) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetTrackOpens returns the TrackOpens field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetTrackOpens() bool {
	if o == nil || IsNil(o.TrackOpens) {
		var ret bool
		return ret
	}
	return *o.TrackOpens
}

// GetTrackOpensOk returns a tuple with the TrackOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetTrackOpensOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackOpens) {
		return nil, false
	}
	return o.TrackOpens, true
}

// HasTrackOpens returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasTrackOpens() bool {
	if o != nil && !IsNil(o.TrackOpens) {
		return true
	}

	return false
}

// SetTrackOpens gets a reference to the given bool and assigns it to the TrackOpens field.
func (o *EmailMessageWithTemplate) SetTrackOpens(v bool) {
	o.TrackOpens = &v
}

// GetTrackClicks returns the TrackClicks field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetTrackClicks() bool {
	if o == nil || IsNil(o.TrackClicks) {
		var ret bool
		return ret
	}
	return *o.TrackClicks
}

// GetTrackClicksOk returns a tuple with the TrackClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetTrackClicksOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackClicks) {
		return nil, false
	}
	return o.TrackClicks, true
}

// HasTrackClicks returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasTrackClicks() bool {
	if o != nil && !IsNil(o.TrackClicks) {
		return true
	}

	return false
}

// SetTrackClicks gets a reference to the given bool and assigns it to the TrackClicks field.
func (o *EmailMessageWithTemplate) SetTrackClicks(v bool) {
	o.TrackClicks = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *EmailMessageWithTemplate) SetGroups(v []string) {
	o.Groups = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *EmailMessageWithTemplate) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetWebhookEndpoint returns the WebhookEndpoint field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetWebhookEndpoint() string {
	if o == nil || IsNil(o.WebhookEndpoint) {
		var ret string
		return ret
	}
	return *o.WebhookEndpoint
}

// GetWebhookEndpointOk returns a tuple with the WebhookEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetWebhookEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEndpoint) {
		return nil, false
	}
	return o.WebhookEndpoint, true
}

// HasWebhookEndpoint returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasWebhookEndpoint() bool {
	if o != nil && !IsNil(o.WebhookEndpoint) {
		return true
	}

	return false
}

// SetWebhookEndpoint gets a reference to the given string and assigns it to the WebhookEndpoint field.
func (o *EmailMessageWithTemplate) SetWebhookEndpoint(v string) {
	o.WebhookEndpoint = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *EmailMessageWithTemplate) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *EmailMessageWithTemplate) GetTemplateVariables() map[string]string {
	if o == nil || IsNil(o.TemplateVariables) {
		var ret map[string]string
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailMessageWithTemplate) GetTemplateVariablesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.TemplateVariables) {
		return map[string]string{}, false
	}
	return o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *EmailMessageWithTemplate) HasTemplateVariables() bool {
	if o != nil && !IsNil(o.TemplateVariables) {
		return true
	}

	return false
}

// SetTemplateVariables gets a reference to the given map[string]string and assigns it to the TemplateVariables field.
func (o *EmailMessageWithTemplate) SetTemplateVariables(v map[string]string) {
	o.TemplateVariables = v
}

func (o EmailMessageWithTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailMessageWithTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	toSerialize["to"] = o.To
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.PreText) {
		toSerialize["preText"] = o.PreText
	}
	if !IsNil(o.HtmlBody) {
		toSerialize["htmlBody"] = o.HtmlBody
	}
	if !IsNil(o.TextBody) {
		toSerialize["textBody"] = o.TextBody
	}
	if !IsNil(o.AmpBody) {
		toSerialize["ampBody"] = o.AmpBody
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Ippool) {
		toSerialize["ippool"] = o.Ippool
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.TrackOpens) {
		toSerialize["trackOpens"] = o.TrackOpens
	}
	if !IsNil(o.TrackClicks) {
		toSerialize["trackClicks"] = o.TrackClicks
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.WebhookEndpoint) {
		toSerialize["webhookEndpoint"] = o.WebhookEndpoint
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.TemplateVariables) {
		toSerialize["templateVariables"] = o.TemplateVariables
	}
	return toSerialize, nil
}

func (o *EmailMessageWithTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"to",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEmailMessageWithTemplate := _EmailMessageWithTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailMessageWithTemplate)

	if err != nil {
		return err
	}

	*o = EmailMessageWithTemplate(varEmailMessageWithTemplate)

	return err
}

type NullableEmailMessageWithTemplate struct {
	value *EmailMessageWithTemplate
	isSet bool
}

func (v NullableEmailMessageWithTemplate) Get() *EmailMessageWithTemplate {
	return v.value
}

func (v *NullableEmailMessageWithTemplate) Set(val *EmailMessageWithTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailMessageWithTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailMessageWithTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailMessageWithTemplate(val *EmailMessageWithTemplate) *NullableEmailMessageWithTemplate {
	return &NullableEmailMessageWithTemplate{value: val, isSet: true}
}

func (v NullableEmailMessageWithTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailMessageWithTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


