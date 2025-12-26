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

// checks if the Webhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Webhook{}

// Webhook struct for Webhook
type Webhook struct {
	// Unique ID for the webhook.
	Id *int32 `json:"id,omitempty"`
	// Indicates if the webhook is active or paused.
	Enabled *bool `json:"enabled,omitempty"`
	// URL endpoint to which webhook calls need to be made.
	Url *string `json:"url,omitempty"`
	// Trigger webhook on email message being processed.
	Processed *bool `json:"processed,omitempty"`
	// Trigger webhook on email message being delivered.
	Delivered *bool `json:"delivered,omitempty"`
	// Trigger webhook on email message being dropped.
	Dropped *bool `json:"dropped,omitempty"`
	// Trigger webhook on email message being soft bounced.
	SoftBounced *bool `json:"softBounced,omitempty"`
	// Trigger webhook on email message being hard bounced.
	HardBounced *bool `json:"hardBounced,omitempty"`
	// Trigger webhook on email message being opened.
	Opened *bool `json:"opened,omitempty"`
	// Trigger webhook on email message link being clicked.
	Clicked *bool `json:"clicked,omitempty"`
	// Trigger webhook on email message being unsubscribed.
	Unsubscribed *bool `json:"unsubscribed,omitempty"`
	// Trigger webhook on email message being marked as spam.
	Spam *bool `json:"spam,omitempty"`
	// Trigger webhook on email message being sent.
	Sent *bool `json:"sent,omitempty"`
	// Trigger webhook on email message being dropped by SMTP.
	SmtpDropped *bool `json:"smtpDropped,omitempty"`
	// Trigger webhook on unique email opens.
	UniqueOpen *bool `json:"uniqueOpen,omitempty"`
	// Trigger webhook on unique email clicks.
	UniqueClick *bool `json:"uniqueClick,omitempty"`
	// UNIX epoch nano timestamp when the webhook was created.
	Created *int64 `json:"created,omitempty"`
	// Member who created the webhook
	CreatedBy map[string]interface{} `json:"created_by,omitempty"`
	// Member who updated the webhook
	UpdatedBy map[string]interface{} `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Webhook Webhook

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook() *Webhook {
	this := Webhook{}
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Webhook) SetId(v int32) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Webhook) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Webhook) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Webhook) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Webhook) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Webhook) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Webhook) SetUrl(v string) {
	o.Url = &v
}

// GetProcessed returns the Processed field value if set, zero value otherwise.
func (o *Webhook) GetProcessed() bool {
	if o == nil || IsNil(o.Processed) {
		var ret bool
		return ret
	}
	return *o.Processed
}

// GetProcessedOk returns a tuple with the Processed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetProcessedOk() (*bool, bool) {
	if o == nil || IsNil(o.Processed) {
		return nil, false
	}
	return o.Processed, true
}

// HasProcessed returns a boolean if a field has been set.
func (o *Webhook) HasProcessed() bool {
	if o != nil && !IsNil(o.Processed) {
		return true
	}

	return false
}

// SetProcessed gets a reference to the given bool and assigns it to the Processed field.
func (o *Webhook) SetProcessed(v bool) {
	o.Processed = &v
}

// GetDelivered returns the Delivered field value if set, zero value otherwise.
func (o *Webhook) GetDelivered() bool {
	if o == nil || IsNil(o.Delivered) {
		var ret bool
		return ret
	}
	return *o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetDeliveredOk() (*bool, bool) {
	if o == nil || IsNil(o.Delivered) {
		return nil, false
	}
	return o.Delivered, true
}

// HasDelivered returns a boolean if a field has been set.
func (o *Webhook) HasDelivered() bool {
	if o != nil && !IsNil(o.Delivered) {
		return true
	}

	return false
}

// SetDelivered gets a reference to the given bool and assigns it to the Delivered field.
func (o *Webhook) SetDelivered(v bool) {
	o.Delivered = &v
}

// GetDropped returns the Dropped field value if set, zero value otherwise.
func (o *Webhook) GetDropped() bool {
	if o == nil || IsNil(o.Dropped) {
		var ret bool
		return ret
	}
	return *o.Dropped
}

// GetDroppedOk returns a tuple with the Dropped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetDroppedOk() (*bool, bool) {
	if o == nil || IsNil(o.Dropped) {
		return nil, false
	}
	return o.Dropped, true
}

// HasDropped returns a boolean if a field has been set.
func (o *Webhook) HasDropped() bool {
	if o != nil && !IsNil(o.Dropped) {
		return true
	}

	return false
}

// SetDropped gets a reference to the given bool and assigns it to the Dropped field.
func (o *Webhook) SetDropped(v bool) {
	o.Dropped = &v
}

// GetSoftBounced returns the SoftBounced field value if set, zero value otherwise.
func (o *Webhook) GetSoftBounced() bool {
	if o == nil || IsNil(o.SoftBounced) {
		var ret bool
		return ret
	}
	return *o.SoftBounced
}

// GetSoftBouncedOk returns a tuple with the SoftBounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSoftBouncedOk() (*bool, bool) {
	if o == nil || IsNil(o.SoftBounced) {
		return nil, false
	}
	return o.SoftBounced, true
}

// HasSoftBounced returns a boolean if a field has been set.
func (o *Webhook) HasSoftBounced() bool {
	if o != nil && !IsNil(o.SoftBounced) {
		return true
	}

	return false
}

// SetSoftBounced gets a reference to the given bool and assigns it to the SoftBounced field.
func (o *Webhook) SetSoftBounced(v bool) {
	o.SoftBounced = &v
}

// GetHardBounced returns the HardBounced field value if set, zero value otherwise.
func (o *Webhook) GetHardBounced() bool {
	if o == nil || IsNil(o.HardBounced) {
		var ret bool
		return ret
	}
	return *o.HardBounced
}

// GetHardBouncedOk returns a tuple with the HardBounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHardBouncedOk() (*bool, bool) {
	if o == nil || IsNil(o.HardBounced) {
		return nil, false
	}
	return o.HardBounced, true
}

// HasHardBounced returns a boolean if a field has been set.
func (o *Webhook) HasHardBounced() bool {
	if o != nil && !IsNil(o.HardBounced) {
		return true
	}

	return false
}

// SetHardBounced gets a reference to the given bool and assigns it to the HardBounced field.
func (o *Webhook) SetHardBounced(v bool) {
	o.HardBounced = &v
}

// GetOpened returns the Opened field value if set, zero value otherwise.
func (o *Webhook) GetOpened() bool {
	if o == nil || IsNil(o.Opened) {
		var ret bool
		return ret
	}
	return *o.Opened
}

// GetOpenedOk returns a tuple with the Opened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetOpenedOk() (*bool, bool) {
	if o == nil || IsNil(o.Opened) {
		return nil, false
	}
	return o.Opened, true
}

// HasOpened returns a boolean if a field has been set.
func (o *Webhook) HasOpened() bool {
	if o != nil && !IsNil(o.Opened) {
		return true
	}

	return false
}

// SetOpened gets a reference to the given bool and assigns it to the Opened field.
func (o *Webhook) SetOpened(v bool) {
	o.Opened = &v
}

// GetClicked returns the Clicked field value if set, zero value otherwise.
func (o *Webhook) GetClicked() bool {
	if o == nil || IsNil(o.Clicked) {
		var ret bool
		return ret
	}
	return *o.Clicked
}

// GetClickedOk returns a tuple with the Clicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetClickedOk() (*bool, bool) {
	if o == nil || IsNil(o.Clicked) {
		return nil, false
	}
	return o.Clicked, true
}

// HasClicked returns a boolean if a field has been set.
func (o *Webhook) HasClicked() bool {
	if o != nil && !IsNil(o.Clicked) {
		return true
	}

	return false
}

// SetClicked gets a reference to the given bool and assigns it to the Clicked field.
func (o *Webhook) SetClicked(v bool) {
	o.Clicked = &v
}

// GetUnsubscribed returns the Unsubscribed field value if set, zero value otherwise.
func (o *Webhook) GetUnsubscribed() bool {
	if o == nil || IsNil(o.Unsubscribed) {
		var ret bool
		return ret
	}
	return *o.Unsubscribed
}

// GetUnsubscribedOk returns a tuple with the Unsubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUnsubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unsubscribed) {
		return nil, false
	}
	return o.Unsubscribed, true
}

// HasUnsubscribed returns a boolean if a field has been set.
func (o *Webhook) HasUnsubscribed() bool {
	if o != nil && !IsNil(o.Unsubscribed) {
		return true
	}

	return false
}

// SetUnsubscribed gets a reference to the given bool and assigns it to the Unsubscribed field.
func (o *Webhook) SetUnsubscribed(v bool) {
	o.Unsubscribed = &v
}

// GetSpam returns the Spam field value if set, zero value otherwise.
func (o *Webhook) GetSpam() bool {
	if o == nil || IsNil(o.Spam) {
		var ret bool
		return ret
	}
	return *o.Spam
}

// GetSpamOk returns a tuple with the Spam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSpamOk() (*bool, bool) {
	if o == nil || IsNil(o.Spam) {
		return nil, false
	}
	return o.Spam, true
}

// HasSpam returns a boolean if a field has been set.
func (o *Webhook) HasSpam() bool {
	if o != nil && !IsNil(o.Spam) {
		return true
	}

	return false
}

// SetSpam gets a reference to the given bool and assigns it to the Spam field.
func (o *Webhook) SetSpam(v bool) {
	o.Spam = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *Webhook) GetSent() bool {
	if o == nil || IsNil(o.Sent) {
		var ret bool
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSentOk() (*bool, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *Webhook) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given bool and assigns it to the Sent field.
func (o *Webhook) SetSent(v bool) {
	o.Sent = &v
}

// GetSmtpDropped returns the SmtpDropped field value if set, zero value otherwise.
func (o *Webhook) GetSmtpDropped() bool {
	if o == nil || IsNil(o.SmtpDropped) {
		var ret bool
		return ret
	}
	return *o.SmtpDropped
}

// GetSmtpDroppedOk returns a tuple with the SmtpDropped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSmtpDroppedOk() (*bool, bool) {
	if o == nil || IsNil(o.SmtpDropped) {
		return nil, false
	}
	return o.SmtpDropped, true
}

// HasSmtpDropped returns a boolean if a field has been set.
func (o *Webhook) HasSmtpDropped() bool {
	if o != nil && !IsNil(o.SmtpDropped) {
		return true
	}

	return false
}

// SetSmtpDropped gets a reference to the given bool and assigns it to the SmtpDropped field.
func (o *Webhook) SetSmtpDropped(v bool) {
	o.SmtpDropped = &v
}

// GetUniqueOpen returns the UniqueOpen field value if set, zero value otherwise.
func (o *Webhook) GetUniqueOpen() bool {
	if o == nil || IsNil(o.UniqueOpen) {
		var ret bool
		return ret
	}
	return *o.UniqueOpen
}

// GetUniqueOpenOk returns a tuple with the UniqueOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUniqueOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueOpen) {
		return nil, false
	}
	return o.UniqueOpen, true
}

// HasUniqueOpen returns a boolean if a field has been set.
func (o *Webhook) HasUniqueOpen() bool {
	if o != nil && !IsNil(o.UniqueOpen) {
		return true
	}

	return false
}

// SetUniqueOpen gets a reference to the given bool and assigns it to the UniqueOpen field.
func (o *Webhook) SetUniqueOpen(v bool) {
	o.UniqueOpen = &v
}

// GetUniqueClick returns the UniqueClick field value if set, zero value otherwise.
func (o *Webhook) GetUniqueClick() bool {
	if o == nil || IsNil(o.UniqueClick) {
		var ret bool
		return ret
	}
	return *o.UniqueClick
}

// GetUniqueClickOk returns a tuple with the UniqueClick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUniqueClickOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueClick) {
		return nil, false
	}
	return o.UniqueClick, true
}

// HasUniqueClick returns a boolean if a field has been set.
func (o *Webhook) HasUniqueClick() bool {
	if o != nil && !IsNil(o.UniqueClick) {
		return true
	}

	return false
}

// SetUniqueClick gets a reference to the given bool and assigns it to the UniqueClick field.
func (o *Webhook) SetUniqueClick(v bool) {
	o.UniqueClick = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Webhook) GetCreated() int64 {
	if o == nil || IsNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Webhook) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Webhook) SetCreated(v int64) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Webhook) GetCreatedBy() map[string]interface{} {
	if o == nil || IsNil(o.CreatedBy) {
		var ret map[string]interface{}
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCreatedByOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return map[string]interface{}{}, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Webhook) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given map[string]interface{} and assigns it to the CreatedBy field.
func (o *Webhook) SetCreatedBy(v map[string]interface{}) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Webhook) GetUpdatedBy() map[string]interface{} {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret map[string]interface{}
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUpdatedByOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return map[string]interface{}{}, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Webhook) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given map[string]interface{} and assigns it to the UpdatedBy field.
func (o *Webhook) SetUpdatedBy(v map[string]interface{}) {
	o.UpdatedBy = v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Processed) {
		toSerialize["processed"] = o.Processed
	}
	if !IsNil(o.Delivered) {
		toSerialize["delivered"] = o.Delivered
	}
	if !IsNil(o.Dropped) {
		toSerialize["dropped"] = o.Dropped
	}
	if !IsNil(o.SoftBounced) {
		toSerialize["softBounced"] = o.SoftBounced
	}
	if !IsNil(o.HardBounced) {
		toSerialize["hardBounced"] = o.HardBounced
	}
	if !IsNil(o.Opened) {
		toSerialize["opened"] = o.Opened
	}
	if !IsNil(o.Clicked) {
		toSerialize["clicked"] = o.Clicked
	}
	if !IsNil(o.Unsubscribed) {
		toSerialize["unsubscribed"] = o.Unsubscribed
	}
	if !IsNil(o.Spam) {
		toSerialize["spam"] = o.Spam
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.SmtpDropped) {
		toSerialize["smtpDropped"] = o.SmtpDropped
	}
	if !IsNil(o.UniqueOpen) {
		toSerialize["uniqueOpen"] = o.UniqueOpen
	}
	if !IsNil(o.UniqueClick) {
		toSerialize["uniqueClick"] = o.UniqueClick
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Webhook) UnmarshalJSON(data []byte) (err error) {
	varWebhook := _Webhook{}

	err = json.Unmarshal(data, &varWebhook)

	if err != nil {
		return err
	}

	*o = Webhook(varWebhook)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "url")
		delete(additionalProperties, "processed")
		delete(additionalProperties, "delivered")
		delete(additionalProperties, "dropped")
		delete(additionalProperties, "softBounced")
		delete(additionalProperties, "hardBounced")
		delete(additionalProperties, "opened")
		delete(additionalProperties, "clicked")
		delete(additionalProperties, "unsubscribed")
		delete(additionalProperties, "spam")
		delete(additionalProperties, "sent")
		delete(additionalProperties, "smtpDropped")
		delete(additionalProperties, "uniqueOpen")
		delete(additionalProperties, "uniqueClick")
		delete(additionalProperties, "created")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


