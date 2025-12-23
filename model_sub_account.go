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

// checks if the SubAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubAccount{}

// SubAccount struct for SubAccount
type SubAccount struct {
	// Unique ID for the sub-account.
	Id *int32 `json:"id,omitempty"`
	// API key for the sub-account.
	ApiKey *string `json:"apiKey,omitempty"`
	// Name of the sub-account.
	Name *string `json:"name,omitempty"`
	// Labels associated with the sub-account
	Labels []string `json:"labels,omitempty"`
	// SMTP Auths associated with the sub-account
	SmtpAuths []SMTPAuth `json:"smtpAuths,omitempty"`
	// Type of the sub-account
	Type *int32 `json:"type,omitempty"`
	// Indicates whether the sub-account is a Plus sub-account
	IsPlus *bool `json:"isPlus,omitempty"`
	// UNIX epoch nano timestamp when the sub-account was created.
	Created *int64 `json:"created,omitempty"`
	// Member who created the sub-account
	CreatedBy Member `json:"created_by,omitempty"`
	// Member who updated the sub-account
	UpdatedBy Member `json:"updated_by,omitempty"`
	// Indicates whether the sub-account is blocked
	Blocked *bool `json:"blocked,omitempty"`
	// UNIX epoch nano timestamp when the sub-account was blocked (0 if not blocked)
	BlockedAt *int32 `json:"blocked_at,omitempty"`
	// Reason for blocking the sub-account
	BlockReason *string `json:"block_reason,omitempty"`
	// Indicates whether the sub-account is exempt from hard bounce tracking
	HbExempt *bool `json:"hb_exempt,omitempty"`
	// Indicates whether weekly reports are generated for this sub-account
	GenerateWeeklyReport *bool `json:"generate_weekly_report,omitempty"`
	// Handlers associated with the sub-account
	Handlers []string `json:"handlers,omitempty"`
}

// NewSubAccount instantiates a new SubAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccount() *SubAccount {
	this := SubAccount{}
	return &this
}

// NewSubAccountWithDefaults instantiates a new SubAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountWithDefaults() *SubAccount {
	this := SubAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubAccount) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SubAccount) SetId(v int32) {
	o.Id = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SubAccount) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SubAccount) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SubAccount) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubAccount) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SubAccount) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SubAccount) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *SubAccount) SetLabels(v []string) {
	o.Labels = v
}

// GetSmtpAuths returns the SmtpAuths field value if set, zero value otherwise.
func (o *SubAccount) GetSmtpAuths() []SMTPAuth {
	if o == nil || IsNil(o.SmtpAuths) {
		var ret []SMTPAuth
		return ret
	}
	return o.SmtpAuths
}

// GetSmtpAuthsOk returns a tuple with the SmtpAuths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetSmtpAuthsOk() ([]SMTPAuth, bool) {
	if o == nil || IsNil(o.SmtpAuths) {
		return nil, false
	}
	return o.SmtpAuths, true
}

// HasSmtpAuths returns a boolean if a field has been set.
func (o *SubAccount) HasSmtpAuths() bool {
	if o != nil && !IsNil(o.SmtpAuths) {
		return true
	}

	return false
}

// SetSmtpAuths gets a reference to the given []SMTPAuth and assigns it to the SmtpAuths field.
func (o *SubAccount) SetSmtpAuths(v []SMTPAuth) {
	o.SmtpAuths = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubAccount) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubAccount) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *SubAccount) SetType(v int32) {
	o.Type = &v
}

// GetIsPlus returns the IsPlus field value if set, zero value otherwise.
func (o *SubAccount) GetIsPlus() bool {
	if o == nil || IsNil(o.IsPlus) {
		var ret bool
		return ret
	}
	return *o.IsPlus
}

// GetIsPlusOk returns a tuple with the IsPlus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetIsPlusOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlus) {
		return nil, false
	}
	return o.IsPlus, true
}

// HasIsPlus returns a boolean if a field has been set.
func (o *SubAccount) HasIsPlus() bool {
	if o != nil && !IsNil(o.IsPlus) {
		return true
	}

	return false
}

// SetIsPlus gets a reference to the given bool and assigns it to the IsPlus field.
func (o *SubAccount) SetIsPlus(v bool) {
	o.IsPlus = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SubAccount) GetCreated() int64 {
	if o == nil || IsNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SubAccount) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *SubAccount) SetCreated(v int64) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SubAccount) GetCreatedBy() Member {
	if o == nil || IsNil(o.CreatedBy) {
		var ret Member
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetCreatedByOk() (Member, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return Member{}, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SubAccount) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given Member and assigns it to the CreatedBy field.
func (o *SubAccount) SetCreatedBy(v Member) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SubAccount) GetUpdatedBy() Member {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret Member
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetUpdatedByOk() (Member, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return Member{}, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SubAccount) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given Member and assigns it to the UpdatedBy field.
func (o *SubAccount) SetUpdatedBy(v Member) {
	o.UpdatedBy = v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *SubAccount) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *SubAccount) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *SubAccount) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetBlockedAt returns the BlockedAt field value if set, zero value otherwise.
func (o *SubAccount) GetBlockedAt() int32 {
	if o == nil || IsNil(o.BlockedAt) {
		var ret int32
		return ret
	}
	return *o.BlockedAt
}

// GetBlockedAtOk returns a tuple with the BlockedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetBlockedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockedAt) {
		return nil, false
	}
	return o.BlockedAt, true
}

// HasBlockedAt returns a boolean if a field has been set.
func (o *SubAccount) HasBlockedAt() bool {
	if o != nil && !IsNil(o.BlockedAt) {
		return true
	}

	return false
}

// SetBlockedAt gets a reference to the given int32 and assigns it to the BlockedAt field.
func (o *SubAccount) SetBlockedAt(v int32) {
	o.BlockedAt = &v
}

// GetBlockReason returns the BlockReason field value if set, zero value otherwise.
func (o *SubAccount) GetBlockReason() string {
	if o == nil || IsNil(o.BlockReason) {
		var ret string
		return ret
	}
	return *o.BlockReason
}

// GetBlockReasonOk returns a tuple with the BlockReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetBlockReasonOk() (*string, bool) {
	if o == nil || IsNil(o.BlockReason) {
		return nil, false
	}
	return o.BlockReason, true
}

// HasBlockReason returns a boolean if a field has been set.
func (o *SubAccount) HasBlockReason() bool {
	if o != nil && !IsNil(o.BlockReason) {
		return true
	}

	return false
}

// SetBlockReason gets a reference to the given string and assigns it to the BlockReason field.
func (o *SubAccount) SetBlockReason(v string) {
	o.BlockReason = &v
}

// GetHbExempt returns the HbExempt field value if set, zero value otherwise.
func (o *SubAccount) GetHbExempt() bool {
	if o == nil || IsNil(o.HbExempt) {
		var ret bool
		return ret
	}
	return *o.HbExempt
}

// GetHbExemptOk returns a tuple with the HbExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetHbExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.HbExempt) {
		return nil, false
	}
	return o.HbExempt, true
}

// HasHbExempt returns a boolean if a field has been set.
func (o *SubAccount) HasHbExempt() bool {
	if o != nil && !IsNil(o.HbExempt) {
		return true
	}

	return false
}

// SetHbExempt gets a reference to the given bool and assigns it to the HbExempt field.
func (o *SubAccount) SetHbExempt(v bool) {
	o.HbExempt = &v
}

// GetGenerateWeeklyReport returns the GenerateWeeklyReport field value if set, zero value otherwise.
func (o *SubAccount) GetGenerateWeeklyReport() bool {
	if o == nil || IsNil(o.GenerateWeeklyReport) {
		var ret bool
		return ret
	}
	return *o.GenerateWeeklyReport
}

// GetGenerateWeeklyReportOk returns a tuple with the GenerateWeeklyReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetGenerateWeeklyReportOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateWeeklyReport) {
		return nil, false
	}
	return o.GenerateWeeklyReport, true
}

// HasGenerateWeeklyReport returns a boolean if a field has been set.
func (o *SubAccount) HasGenerateWeeklyReport() bool {
	if o != nil && !IsNil(o.GenerateWeeklyReport) {
		return true
	}

	return false
}

// SetGenerateWeeklyReport gets a reference to the given bool and assigns it to the GenerateWeeklyReport field.
func (o *SubAccount) SetGenerateWeeklyReport(v bool) {
	o.GenerateWeeklyReport = &v
}

// GetHandlers returns the Handlers field value if set, zero value otherwise.
func (o *SubAccount) GetHandlers() []string {
	if o == nil || IsNil(o.Handlers) {
		var ret []string
		return ret
	}
	return o.Handlers
}

// GetHandlersOk returns a tuple with the Handlers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetHandlersOk() ([]string, bool) {
	if o == nil || IsNil(o.Handlers) {
		return nil, false
	}
	return o.Handlers, true
}

// HasHandlers returns a boolean if a field has been set.
func (o *SubAccount) HasHandlers() bool {
	if o != nil && !IsNil(o.Handlers) {
		return true
	}

	return false
}

// SetHandlers gets a reference to the given []string and assigns it to the Handlers field.
func (o *SubAccount) SetHandlers(v []string) {
	o.Handlers = v
}

func (o SubAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.SmtpAuths) {
		toSerialize["smtpAuths"] = o.SmtpAuths
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsPlus) {
		toSerialize["isPlus"] = o.IsPlus
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
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.BlockedAt) {
		toSerialize["blocked_at"] = o.BlockedAt
	}
	if !IsNil(o.BlockReason) {
		toSerialize["block_reason"] = o.BlockReason
	}
	if !IsNil(o.HbExempt) {
		toSerialize["hb_exempt"] = o.HbExempt
	}
	if !IsNil(o.GenerateWeeklyReport) {
		toSerialize["generate_weekly_report"] = o.GenerateWeeklyReport
	}
	if !IsNil(o.Handlers) {
		toSerialize["handlers"] = o.Handlers
	}
	return toSerialize, nil
}

type NullableSubAccount struct {
	value *SubAccount
	isSet bool
}

func (v NullableSubAccount) Get() *SubAccount {
	return v.value
}

func (v *NullableSubAccount) Set(val *SubAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccount(val *SubAccount) *NullableSubAccount {
	return &NullableSubAccount{value: val, isSet: true}
}

func (v NullableSubAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


