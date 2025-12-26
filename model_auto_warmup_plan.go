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

// checks if the AutoWarmupPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoWarmupPlan{}

// AutoWarmupPlan struct for AutoWarmupPlan
type AutoWarmupPlan struct {
	// Unique ID for the auto-warmup plan
	Id *int32 `json:"id,omitempty"`
	// Name of the auto-warmup plan
	Name *string `json:"name,omitempty"`
	// Gmail warmup plan configuration in JSON format
	GmailWarmupPlan *string `json:"gmailWarmupPlan,omitempty"`
	// Yahoo warmup plan configuration in JSON format
	YahooWarmupPlan *string `json:"yahooWarmupPlan,omitempty"`
	// AOL warmup plan configuration in JSON format
	AolWarmupPlan *string `json:"aolWarmupPlan,omitempty"`
	// Microsoft warmup plan configuration in JSON format
	MicrosoftWarmupPlan *string `json:"microsoftWarmupPlan,omitempty"`
	// Comcast warmup plan configuration in JSON format
	ComcastWarmupPlan *string `json:"comcastWarmupPlan,omitempty"`
	// Yandex warmup plan configuration in JSON format
	YandexWarmupPlan *string `json:"yandexWarmupPlan,omitempty"`
	// GMX warmup plan configuration in JSON format
	GmxWarmupPlan *string `json:"gmxWarmupPlan,omitempty"`
	// Mail.ru warmup plan configuration in JSON format
	MailruWarmupPlan *string `json:"mailruWarmupPlan,omitempty"`
	// iCloud warmup plan configuration in JSON format
	IcloudWarmupPlan *string `json:"icloudWarmupPlan,omitempty"`
	// Zoho warmup plan configuration in JSON format
	ZohoWarmupPlan *string `json:"zohoWarmupPlan,omitempty"`
	// QQ warmup plan configuration in JSON format
	QqWarmupPlan *string `json:"qqWarmupPlan,omitempty"`
	// Default warmup plan configuration in JSON format
	DefaultWarmupPlan *string `json:"defaultWarmupPlan,omitempty"`
	// AT&T warmup plan configuration in JSON format
	AttWarmupPlan *string `json:"attWarmupPlan,omitempty"`
	// Office365 warmup plan configuration in JSON format
	Office365WarmupPlan *string `json:"office365WarmupPlan,omitempty"`
	// Google Workspace warmup plan configuration in JSON format
	GoogleworkspaceWarmupPlan *string `json:"googleworkspaceWarmupPlan,omitempty"`
	// Proofpoint warmup plan configuration in JSON format
	ProofpointWarmupPlan *string `json:"proofpointWarmupPlan,omitempty"`
	// Mimecast warmup plan configuration in JSON format
	MimecastWarmupPlan *string `json:"mimecastWarmupPlan,omitempty"`
	// Barracuda warmup plan configuration in JSON format
	BarracudaWarmupPlan *string `json:"barracudaWarmupPlan,omitempty"`
	// Cisco IronPort warmup plan configuration in JSON format
	CiscoironportWarmupPlan *string `json:"ciscoironportWarmupPlan,omitempty"`
	// Rackspace warmup plan configuration in JSON format
	RackspaceWarmupPlan *string `json:"rackspaceWarmupPlan,omitempty"`
	// Zoho Business warmup plan configuration in JSON format
	ZohobusinessWarmupPlan *string `json:"zohobusinessWarmupPlan,omitempty"`
	// Amazon WorkMail warmup plan configuration in JSON format
	AmazonworkmailWarmupPlan *string `json:"amazonworkmailWarmupPlan,omitempty"`
	// Symantec warmup plan configuration in JSON format
	SymantecWarmupPlan *string `json:"symantecWarmupPlan,omitempty"`
	// Fortinet warmup plan configuration in JSON format
	FortinetWarmupPlan *string `json:"fortinetWarmupPlan,omitempty"`
	// Sophos warmup plan configuration in JSON format
	SophosWarmupPlan *string `json:"sophosWarmupPlan,omitempty"`
	// Trend Micro warmup plan configuration in JSON format
	TrendmicroWarmupPlan *string `json:"trendmicroWarmupPlan,omitempty"`
	// Checkpoint warmup plan configuration in JSON format
	CheckpointWarmupPlan *string `json:"checkpointWarmupPlan,omitempty"`
	// UNIX epoch nano timestamp when the warmup plan was created
	Created *int64 `json:"created,omitempty"`
	// UNIX epoch nano timestamp when the warmup plan was last updated
	Updated *int64 `json:"updated,omitempty"`
	// Warmup interval in hours
	WarmupInterval *int32 `json:"warmupInterval,omitempty"`
}

// NewAutoWarmupPlan instantiates a new AutoWarmupPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoWarmupPlan() *AutoWarmupPlan {
	this := AutoWarmupPlan{}
	return &this
}

// NewAutoWarmupPlanWithDefaults instantiates a new AutoWarmupPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoWarmupPlanWithDefaults() *AutoWarmupPlan {
	this := AutoWarmupPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AutoWarmupPlan) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutoWarmupPlan) SetName(v string) {
	o.Name = &v
}

// GetGmailWarmupPlan returns the GmailWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetGmailWarmupPlan() string {
	if o == nil || IsNil(o.GmailWarmupPlan) {
		var ret string
		return ret
	}
	return *o.GmailWarmupPlan
}

// GetGmailWarmupPlanOk returns a tuple with the GmailWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetGmailWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.GmailWarmupPlan) {
		return nil, false
	}
	return o.GmailWarmupPlan, true
}

// HasGmailWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasGmailWarmupPlan() bool {
	if o != nil && !IsNil(o.GmailWarmupPlan) {
		return true
	}

	return false
}

// SetGmailWarmupPlan gets a reference to the given string and assigns it to the GmailWarmupPlan field.
func (o *AutoWarmupPlan) SetGmailWarmupPlan(v string) {
	o.GmailWarmupPlan = &v
}

// GetYahooWarmupPlan returns the YahooWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetYahooWarmupPlan() string {
	if o == nil || IsNil(o.YahooWarmupPlan) {
		var ret string
		return ret
	}
	return *o.YahooWarmupPlan
}

// GetYahooWarmupPlanOk returns a tuple with the YahooWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetYahooWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.YahooWarmupPlan) {
		return nil, false
	}
	return o.YahooWarmupPlan, true
}

// HasYahooWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasYahooWarmupPlan() bool {
	if o != nil && !IsNil(o.YahooWarmupPlan) {
		return true
	}

	return false
}

// SetYahooWarmupPlan gets a reference to the given string and assigns it to the YahooWarmupPlan field.
func (o *AutoWarmupPlan) SetYahooWarmupPlan(v string) {
	o.YahooWarmupPlan = &v
}

// GetAolWarmupPlan returns the AolWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetAolWarmupPlan() string {
	if o == nil || IsNil(o.AolWarmupPlan) {
		var ret string
		return ret
	}
	return *o.AolWarmupPlan
}

// GetAolWarmupPlanOk returns a tuple with the AolWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetAolWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.AolWarmupPlan) {
		return nil, false
	}
	return o.AolWarmupPlan, true
}

// HasAolWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasAolWarmupPlan() bool {
	if o != nil && !IsNil(o.AolWarmupPlan) {
		return true
	}

	return false
}

// SetAolWarmupPlan gets a reference to the given string and assigns it to the AolWarmupPlan field.
func (o *AutoWarmupPlan) SetAolWarmupPlan(v string) {
	o.AolWarmupPlan = &v
}

// GetMicrosoftWarmupPlan returns the MicrosoftWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetMicrosoftWarmupPlan() string {
	if o == nil || IsNil(o.MicrosoftWarmupPlan) {
		var ret string
		return ret
	}
	return *o.MicrosoftWarmupPlan
}

// GetMicrosoftWarmupPlanOk returns a tuple with the MicrosoftWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetMicrosoftWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.MicrosoftWarmupPlan) {
		return nil, false
	}
	return o.MicrosoftWarmupPlan, true
}

// HasMicrosoftWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasMicrosoftWarmupPlan() bool {
	if o != nil && !IsNil(o.MicrosoftWarmupPlan) {
		return true
	}

	return false
}

// SetMicrosoftWarmupPlan gets a reference to the given string and assigns it to the MicrosoftWarmupPlan field.
func (o *AutoWarmupPlan) SetMicrosoftWarmupPlan(v string) {
	o.MicrosoftWarmupPlan = &v
}

// GetComcastWarmupPlan returns the ComcastWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetComcastWarmupPlan() string {
	if o == nil || IsNil(o.ComcastWarmupPlan) {
		var ret string
		return ret
	}
	return *o.ComcastWarmupPlan
}

// GetComcastWarmupPlanOk returns a tuple with the ComcastWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetComcastWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.ComcastWarmupPlan) {
		return nil, false
	}
	return o.ComcastWarmupPlan, true
}

// HasComcastWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasComcastWarmupPlan() bool {
	if o != nil && !IsNil(o.ComcastWarmupPlan) {
		return true
	}

	return false
}

// SetComcastWarmupPlan gets a reference to the given string and assigns it to the ComcastWarmupPlan field.
func (o *AutoWarmupPlan) SetComcastWarmupPlan(v string) {
	o.ComcastWarmupPlan = &v
}

// GetYandexWarmupPlan returns the YandexWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetYandexWarmupPlan() string {
	if o == nil || IsNil(o.YandexWarmupPlan) {
		var ret string
		return ret
	}
	return *o.YandexWarmupPlan
}

// GetYandexWarmupPlanOk returns a tuple with the YandexWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetYandexWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.YandexWarmupPlan) {
		return nil, false
	}
	return o.YandexWarmupPlan, true
}

// HasYandexWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasYandexWarmupPlan() bool {
	if o != nil && !IsNil(o.YandexWarmupPlan) {
		return true
	}

	return false
}

// SetYandexWarmupPlan gets a reference to the given string and assigns it to the YandexWarmupPlan field.
func (o *AutoWarmupPlan) SetYandexWarmupPlan(v string) {
	o.YandexWarmupPlan = &v
}

// GetGmxWarmupPlan returns the GmxWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetGmxWarmupPlan() string {
	if o == nil || IsNil(o.GmxWarmupPlan) {
		var ret string
		return ret
	}
	return *o.GmxWarmupPlan
}

// GetGmxWarmupPlanOk returns a tuple with the GmxWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetGmxWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.GmxWarmupPlan) {
		return nil, false
	}
	return o.GmxWarmupPlan, true
}

// HasGmxWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasGmxWarmupPlan() bool {
	if o != nil && !IsNil(o.GmxWarmupPlan) {
		return true
	}

	return false
}

// SetGmxWarmupPlan gets a reference to the given string and assigns it to the GmxWarmupPlan field.
func (o *AutoWarmupPlan) SetGmxWarmupPlan(v string) {
	o.GmxWarmupPlan = &v
}

// GetMailruWarmupPlan returns the MailruWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetMailruWarmupPlan() string {
	if o == nil || IsNil(o.MailruWarmupPlan) {
		var ret string
		return ret
	}
	return *o.MailruWarmupPlan
}

// GetMailruWarmupPlanOk returns a tuple with the MailruWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetMailruWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.MailruWarmupPlan) {
		return nil, false
	}
	return o.MailruWarmupPlan, true
}

// HasMailruWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasMailruWarmupPlan() bool {
	if o != nil && !IsNil(o.MailruWarmupPlan) {
		return true
	}

	return false
}

// SetMailruWarmupPlan gets a reference to the given string and assigns it to the MailruWarmupPlan field.
func (o *AutoWarmupPlan) SetMailruWarmupPlan(v string) {
	o.MailruWarmupPlan = &v
}

// GetIcloudWarmupPlan returns the IcloudWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetIcloudWarmupPlan() string {
	if o == nil || IsNil(o.IcloudWarmupPlan) {
		var ret string
		return ret
	}
	return *o.IcloudWarmupPlan
}

// GetIcloudWarmupPlanOk returns a tuple with the IcloudWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetIcloudWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.IcloudWarmupPlan) {
		return nil, false
	}
	return o.IcloudWarmupPlan, true
}

// HasIcloudWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasIcloudWarmupPlan() bool {
	if o != nil && !IsNil(o.IcloudWarmupPlan) {
		return true
	}

	return false
}

// SetIcloudWarmupPlan gets a reference to the given string and assigns it to the IcloudWarmupPlan field.
func (o *AutoWarmupPlan) SetIcloudWarmupPlan(v string) {
	o.IcloudWarmupPlan = &v
}

// GetZohoWarmupPlan returns the ZohoWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetZohoWarmupPlan() string {
	if o == nil || IsNil(o.ZohoWarmupPlan) {
		var ret string
		return ret
	}
	return *o.ZohoWarmupPlan
}

// GetZohoWarmupPlanOk returns a tuple with the ZohoWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetZohoWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.ZohoWarmupPlan) {
		return nil, false
	}
	return o.ZohoWarmupPlan, true
}

// HasZohoWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasZohoWarmupPlan() bool {
	if o != nil && !IsNil(o.ZohoWarmupPlan) {
		return true
	}

	return false
}

// SetZohoWarmupPlan gets a reference to the given string and assigns it to the ZohoWarmupPlan field.
func (o *AutoWarmupPlan) SetZohoWarmupPlan(v string) {
	o.ZohoWarmupPlan = &v
}

// GetQqWarmupPlan returns the QqWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetQqWarmupPlan() string {
	if o == nil || IsNil(o.QqWarmupPlan) {
		var ret string
		return ret
	}
	return *o.QqWarmupPlan
}

// GetQqWarmupPlanOk returns a tuple with the QqWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetQqWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.QqWarmupPlan) {
		return nil, false
	}
	return o.QqWarmupPlan, true
}

// HasQqWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasQqWarmupPlan() bool {
	if o != nil && !IsNil(o.QqWarmupPlan) {
		return true
	}

	return false
}

// SetQqWarmupPlan gets a reference to the given string and assigns it to the QqWarmupPlan field.
func (o *AutoWarmupPlan) SetQqWarmupPlan(v string) {
	o.QqWarmupPlan = &v
}

// GetDefaultWarmupPlan returns the DefaultWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetDefaultWarmupPlan() string {
	if o == nil || IsNil(o.DefaultWarmupPlan) {
		var ret string
		return ret
	}
	return *o.DefaultWarmupPlan
}

// GetDefaultWarmupPlanOk returns a tuple with the DefaultWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetDefaultWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultWarmupPlan) {
		return nil, false
	}
	return o.DefaultWarmupPlan, true
}

// HasDefaultWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasDefaultWarmupPlan() bool {
	if o != nil && !IsNil(o.DefaultWarmupPlan) {
		return true
	}

	return false
}

// SetDefaultWarmupPlan gets a reference to the given string and assigns it to the DefaultWarmupPlan field.
func (o *AutoWarmupPlan) SetDefaultWarmupPlan(v string) {
	o.DefaultWarmupPlan = &v
}

// GetAttWarmupPlan returns the AttWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetAttWarmupPlan() string {
	if o == nil || IsNil(o.AttWarmupPlan) {
		var ret string
		return ret
	}
	return *o.AttWarmupPlan
}

// GetAttWarmupPlanOk returns a tuple with the AttWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetAttWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.AttWarmupPlan) {
		return nil, false
	}
	return o.AttWarmupPlan, true
}

// HasAttWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasAttWarmupPlan() bool {
	if o != nil && !IsNil(o.AttWarmupPlan) {
		return true
	}

	return false
}

// SetAttWarmupPlan gets a reference to the given string and assigns it to the AttWarmupPlan field.
func (o *AutoWarmupPlan) SetAttWarmupPlan(v string) {
	o.AttWarmupPlan = &v
}

// GetOffice365WarmupPlan returns the Office365WarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetOffice365WarmupPlan() string {
	if o == nil || IsNil(o.Office365WarmupPlan) {
		var ret string
		return ret
	}
	return *o.Office365WarmupPlan
}

// GetOffice365WarmupPlanOk returns a tuple with the Office365WarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetOffice365WarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Office365WarmupPlan) {
		return nil, false
	}
	return o.Office365WarmupPlan, true
}

// HasOffice365WarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasOffice365WarmupPlan() bool {
	if o != nil && !IsNil(o.Office365WarmupPlan) {
		return true
	}

	return false
}

// SetOffice365WarmupPlan gets a reference to the given string and assigns it to the Office365WarmupPlan field.
func (o *AutoWarmupPlan) SetOffice365WarmupPlan(v string) {
	o.Office365WarmupPlan = &v
}

// GetGoogleworkspaceWarmupPlan returns the GoogleworkspaceWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetGoogleworkspaceWarmupPlan() string {
	if o == nil || IsNil(o.GoogleworkspaceWarmupPlan) {
		var ret string
		return ret
	}
	return *o.GoogleworkspaceWarmupPlan
}

// GetGoogleworkspaceWarmupPlanOk returns a tuple with the GoogleworkspaceWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetGoogleworkspaceWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleworkspaceWarmupPlan) {
		return nil, false
	}
	return o.GoogleworkspaceWarmupPlan, true
}

// HasGoogleworkspaceWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasGoogleworkspaceWarmupPlan() bool {
	if o != nil && !IsNil(o.GoogleworkspaceWarmupPlan) {
		return true
	}

	return false
}

// SetGoogleworkspaceWarmupPlan gets a reference to the given string and assigns it to the GoogleworkspaceWarmupPlan field.
func (o *AutoWarmupPlan) SetGoogleworkspaceWarmupPlan(v string) {
	o.GoogleworkspaceWarmupPlan = &v
}

// GetProofpointWarmupPlan returns the ProofpointWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetProofpointWarmupPlan() string {
	if o == nil || IsNil(o.ProofpointWarmupPlan) {
		var ret string
		return ret
	}
	return *o.ProofpointWarmupPlan
}

// GetProofpointWarmupPlanOk returns a tuple with the ProofpointWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetProofpointWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.ProofpointWarmupPlan) {
		return nil, false
	}
	return o.ProofpointWarmupPlan, true
}

// HasProofpointWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasProofpointWarmupPlan() bool {
	if o != nil && !IsNil(o.ProofpointWarmupPlan) {
		return true
	}

	return false
}

// SetProofpointWarmupPlan gets a reference to the given string and assigns it to the ProofpointWarmupPlan field.
func (o *AutoWarmupPlan) SetProofpointWarmupPlan(v string) {
	o.ProofpointWarmupPlan = &v
}

// GetMimecastWarmupPlan returns the MimecastWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetMimecastWarmupPlan() string {
	if o == nil || IsNil(o.MimecastWarmupPlan) {
		var ret string
		return ret
	}
	return *o.MimecastWarmupPlan
}

// GetMimecastWarmupPlanOk returns a tuple with the MimecastWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetMimecastWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.MimecastWarmupPlan) {
		return nil, false
	}
	return o.MimecastWarmupPlan, true
}

// HasMimecastWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasMimecastWarmupPlan() bool {
	if o != nil && !IsNil(o.MimecastWarmupPlan) {
		return true
	}

	return false
}

// SetMimecastWarmupPlan gets a reference to the given string and assigns it to the MimecastWarmupPlan field.
func (o *AutoWarmupPlan) SetMimecastWarmupPlan(v string) {
	o.MimecastWarmupPlan = &v
}

// GetBarracudaWarmupPlan returns the BarracudaWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetBarracudaWarmupPlan() string {
	if o == nil || IsNil(o.BarracudaWarmupPlan) {
		var ret string
		return ret
	}
	return *o.BarracudaWarmupPlan
}

// GetBarracudaWarmupPlanOk returns a tuple with the BarracudaWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetBarracudaWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.BarracudaWarmupPlan) {
		return nil, false
	}
	return o.BarracudaWarmupPlan, true
}

// HasBarracudaWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasBarracudaWarmupPlan() bool {
	if o != nil && !IsNil(o.BarracudaWarmupPlan) {
		return true
	}

	return false
}

// SetBarracudaWarmupPlan gets a reference to the given string and assigns it to the BarracudaWarmupPlan field.
func (o *AutoWarmupPlan) SetBarracudaWarmupPlan(v string) {
	o.BarracudaWarmupPlan = &v
}

// GetCiscoironportWarmupPlan returns the CiscoironportWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetCiscoironportWarmupPlan() string {
	if o == nil || IsNil(o.CiscoironportWarmupPlan) {
		var ret string
		return ret
	}
	return *o.CiscoironportWarmupPlan
}

// GetCiscoironportWarmupPlanOk returns a tuple with the CiscoironportWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetCiscoironportWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.CiscoironportWarmupPlan) {
		return nil, false
	}
	return o.CiscoironportWarmupPlan, true
}

// HasCiscoironportWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasCiscoironportWarmupPlan() bool {
	if o != nil && !IsNil(o.CiscoironportWarmupPlan) {
		return true
	}

	return false
}

// SetCiscoironportWarmupPlan gets a reference to the given string and assigns it to the CiscoironportWarmupPlan field.
func (o *AutoWarmupPlan) SetCiscoironportWarmupPlan(v string) {
	o.CiscoironportWarmupPlan = &v
}

// GetRackspaceWarmupPlan returns the RackspaceWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetRackspaceWarmupPlan() string {
	if o == nil || IsNil(o.RackspaceWarmupPlan) {
		var ret string
		return ret
	}
	return *o.RackspaceWarmupPlan
}

// GetRackspaceWarmupPlanOk returns a tuple with the RackspaceWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetRackspaceWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.RackspaceWarmupPlan) {
		return nil, false
	}
	return o.RackspaceWarmupPlan, true
}

// HasRackspaceWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasRackspaceWarmupPlan() bool {
	if o != nil && !IsNil(o.RackspaceWarmupPlan) {
		return true
	}

	return false
}

// SetRackspaceWarmupPlan gets a reference to the given string and assigns it to the RackspaceWarmupPlan field.
func (o *AutoWarmupPlan) SetRackspaceWarmupPlan(v string) {
	o.RackspaceWarmupPlan = &v
}

// GetZohobusinessWarmupPlan returns the ZohobusinessWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetZohobusinessWarmupPlan() string {
	if o == nil || IsNil(o.ZohobusinessWarmupPlan) {
		var ret string
		return ret
	}
	return *o.ZohobusinessWarmupPlan
}

// GetZohobusinessWarmupPlanOk returns a tuple with the ZohobusinessWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetZohobusinessWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.ZohobusinessWarmupPlan) {
		return nil, false
	}
	return o.ZohobusinessWarmupPlan, true
}

// HasZohobusinessWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasZohobusinessWarmupPlan() bool {
	if o != nil && !IsNil(o.ZohobusinessWarmupPlan) {
		return true
	}

	return false
}

// SetZohobusinessWarmupPlan gets a reference to the given string and assigns it to the ZohobusinessWarmupPlan field.
func (o *AutoWarmupPlan) SetZohobusinessWarmupPlan(v string) {
	o.ZohobusinessWarmupPlan = &v
}

// GetAmazonworkmailWarmupPlan returns the AmazonworkmailWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetAmazonworkmailWarmupPlan() string {
	if o == nil || IsNil(o.AmazonworkmailWarmupPlan) {
		var ret string
		return ret
	}
	return *o.AmazonworkmailWarmupPlan
}

// GetAmazonworkmailWarmupPlanOk returns a tuple with the AmazonworkmailWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetAmazonworkmailWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonworkmailWarmupPlan) {
		return nil, false
	}
	return o.AmazonworkmailWarmupPlan, true
}

// HasAmazonworkmailWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasAmazonworkmailWarmupPlan() bool {
	if o != nil && !IsNil(o.AmazonworkmailWarmupPlan) {
		return true
	}

	return false
}

// SetAmazonworkmailWarmupPlan gets a reference to the given string and assigns it to the AmazonworkmailWarmupPlan field.
func (o *AutoWarmupPlan) SetAmazonworkmailWarmupPlan(v string) {
	o.AmazonworkmailWarmupPlan = &v
}

// GetSymantecWarmupPlan returns the SymantecWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetSymantecWarmupPlan() string {
	if o == nil || IsNil(o.SymantecWarmupPlan) {
		var ret string
		return ret
	}
	return *o.SymantecWarmupPlan
}

// GetSymantecWarmupPlanOk returns a tuple with the SymantecWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetSymantecWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.SymantecWarmupPlan) {
		return nil, false
	}
	return o.SymantecWarmupPlan, true
}

// HasSymantecWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasSymantecWarmupPlan() bool {
	if o != nil && !IsNil(o.SymantecWarmupPlan) {
		return true
	}

	return false
}

// SetSymantecWarmupPlan gets a reference to the given string and assigns it to the SymantecWarmupPlan field.
func (o *AutoWarmupPlan) SetSymantecWarmupPlan(v string) {
	o.SymantecWarmupPlan = &v
}

// GetFortinetWarmupPlan returns the FortinetWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetFortinetWarmupPlan() string {
	if o == nil || IsNil(o.FortinetWarmupPlan) {
		var ret string
		return ret
	}
	return *o.FortinetWarmupPlan
}

// GetFortinetWarmupPlanOk returns a tuple with the FortinetWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetFortinetWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.FortinetWarmupPlan) {
		return nil, false
	}
	return o.FortinetWarmupPlan, true
}

// HasFortinetWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasFortinetWarmupPlan() bool {
	if o != nil && !IsNil(o.FortinetWarmupPlan) {
		return true
	}

	return false
}

// SetFortinetWarmupPlan gets a reference to the given string and assigns it to the FortinetWarmupPlan field.
func (o *AutoWarmupPlan) SetFortinetWarmupPlan(v string) {
	o.FortinetWarmupPlan = &v
}

// GetSophosWarmupPlan returns the SophosWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetSophosWarmupPlan() string {
	if o == nil || IsNil(o.SophosWarmupPlan) {
		var ret string
		return ret
	}
	return *o.SophosWarmupPlan
}

// GetSophosWarmupPlanOk returns a tuple with the SophosWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetSophosWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.SophosWarmupPlan) {
		return nil, false
	}
	return o.SophosWarmupPlan, true
}

// HasSophosWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasSophosWarmupPlan() bool {
	if o != nil && !IsNil(o.SophosWarmupPlan) {
		return true
	}

	return false
}

// SetSophosWarmupPlan gets a reference to the given string and assigns it to the SophosWarmupPlan field.
func (o *AutoWarmupPlan) SetSophosWarmupPlan(v string) {
	o.SophosWarmupPlan = &v
}

// GetTrendmicroWarmupPlan returns the TrendmicroWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetTrendmicroWarmupPlan() string {
	if o == nil || IsNil(o.TrendmicroWarmupPlan) {
		var ret string
		return ret
	}
	return *o.TrendmicroWarmupPlan
}

// GetTrendmicroWarmupPlanOk returns a tuple with the TrendmicroWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetTrendmicroWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.TrendmicroWarmupPlan) {
		return nil, false
	}
	return o.TrendmicroWarmupPlan, true
}

// HasTrendmicroWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasTrendmicroWarmupPlan() bool {
	if o != nil && !IsNil(o.TrendmicroWarmupPlan) {
		return true
	}

	return false
}

// SetTrendmicroWarmupPlan gets a reference to the given string and assigns it to the TrendmicroWarmupPlan field.
func (o *AutoWarmupPlan) SetTrendmicroWarmupPlan(v string) {
	o.TrendmicroWarmupPlan = &v
}

// GetCheckpointWarmupPlan returns the CheckpointWarmupPlan field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetCheckpointWarmupPlan() string {
	if o == nil || IsNil(o.CheckpointWarmupPlan) {
		var ret string
		return ret
	}
	return *o.CheckpointWarmupPlan
}

// GetCheckpointWarmupPlanOk returns a tuple with the CheckpointWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetCheckpointWarmupPlanOk() (*string, bool) {
	if o == nil || IsNil(o.CheckpointWarmupPlan) {
		return nil, false
	}
	return o.CheckpointWarmupPlan, true
}

// HasCheckpointWarmupPlan returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasCheckpointWarmupPlan() bool {
	if o != nil && !IsNil(o.CheckpointWarmupPlan) {
		return true
	}

	return false
}

// SetCheckpointWarmupPlan gets a reference to the given string and assigns it to the CheckpointWarmupPlan field.
func (o *AutoWarmupPlan) SetCheckpointWarmupPlan(v string) {
	o.CheckpointWarmupPlan = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetCreated() int64 {
	if o == nil || IsNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *AutoWarmupPlan) SetCreated(v int64) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetUpdated() int64 {
	if o == nil || IsNil(o.Updated) {
		var ret int64
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetUpdatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given int64 and assigns it to the Updated field.
func (o *AutoWarmupPlan) SetUpdated(v int64) {
	o.Updated = &v
}

// GetWarmupInterval returns the WarmupInterval field value if set, zero value otherwise.
func (o *AutoWarmupPlan) GetWarmupInterval() int32 {
	if o == nil || IsNil(o.WarmupInterval) {
		var ret int32
		return ret
	}
	return *o.WarmupInterval
}

// GetWarmupIntervalOk returns a tuple with the WarmupInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoWarmupPlan) GetWarmupIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.WarmupInterval) {
		return nil, false
	}
	return o.WarmupInterval, true
}

// HasWarmupInterval returns a boolean if a field has been set.
func (o *AutoWarmupPlan) HasWarmupInterval() bool {
	if o != nil && !IsNil(o.WarmupInterval) {
		return true
	}

	return false
}

// SetWarmupInterval gets a reference to the given int32 and assigns it to the WarmupInterval field.
func (o *AutoWarmupPlan) SetWarmupInterval(v int32) {
	o.WarmupInterval = &v
}

func (o AutoWarmupPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoWarmupPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.GmailWarmupPlan) {
		toSerialize["gmailWarmupPlan"] = o.GmailWarmupPlan
	}
	if !IsNil(o.YahooWarmupPlan) {
		toSerialize["yahooWarmupPlan"] = o.YahooWarmupPlan
	}
	if !IsNil(o.AolWarmupPlan) {
		toSerialize["aolWarmupPlan"] = o.AolWarmupPlan
	}
	if !IsNil(o.MicrosoftWarmupPlan) {
		toSerialize["microsoftWarmupPlan"] = o.MicrosoftWarmupPlan
	}
	if !IsNil(o.ComcastWarmupPlan) {
		toSerialize["comcastWarmupPlan"] = o.ComcastWarmupPlan
	}
	if !IsNil(o.YandexWarmupPlan) {
		toSerialize["yandexWarmupPlan"] = o.YandexWarmupPlan
	}
	if !IsNil(o.GmxWarmupPlan) {
		toSerialize["gmxWarmupPlan"] = o.GmxWarmupPlan
	}
	if !IsNil(o.MailruWarmupPlan) {
		toSerialize["mailruWarmupPlan"] = o.MailruWarmupPlan
	}
	if !IsNil(o.IcloudWarmupPlan) {
		toSerialize["icloudWarmupPlan"] = o.IcloudWarmupPlan
	}
	if !IsNil(o.ZohoWarmupPlan) {
		toSerialize["zohoWarmupPlan"] = o.ZohoWarmupPlan
	}
	if !IsNil(o.QqWarmupPlan) {
		toSerialize["qqWarmupPlan"] = o.QqWarmupPlan
	}
	if !IsNil(o.DefaultWarmupPlan) {
		toSerialize["defaultWarmupPlan"] = o.DefaultWarmupPlan
	}
	if !IsNil(o.AttWarmupPlan) {
		toSerialize["attWarmupPlan"] = o.AttWarmupPlan
	}
	if !IsNil(o.Office365WarmupPlan) {
		toSerialize["office365WarmupPlan"] = o.Office365WarmupPlan
	}
	if !IsNil(o.GoogleworkspaceWarmupPlan) {
		toSerialize["googleworkspaceWarmupPlan"] = o.GoogleworkspaceWarmupPlan
	}
	if !IsNil(o.ProofpointWarmupPlan) {
		toSerialize["proofpointWarmupPlan"] = o.ProofpointWarmupPlan
	}
	if !IsNil(o.MimecastWarmupPlan) {
		toSerialize["mimecastWarmupPlan"] = o.MimecastWarmupPlan
	}
	if !IsNil(o.BarracudaWarmupPlan) {
		toSerialize["barracudaWarmupPlan"] = o.BarracudaWarmupPlan
	}
	if !IsNil(o.CiscoironportWarmupPlan) {
		toSerialize["ciscoironportWarmupPlan"] = o.CiscoironportWarmupPlan
	}
	if !IsNil(o.RackspaceWarmupPlan) {
		toSerialize["rackspaceWarmupPlan"] = o.RackspaceWarmupPlan
	}
	if !IsNil(o.ZohobusinessWarmupPlan) {
		toSerialize["zohobusinessWarmupPlan"] = o.ZohobusinessWarmupPlan
	}
	if !IsNil(o.AmazonworkmailWarmupPlan) {
		toSerialize["amazonworkmailWarmupPlan"] = o.AmazonworkmailWarmupPlan
	}
	if !IsNil(o.SymantecWarmupPlan) {
		toSerialize["symantecWarmupPlan"] = o.SymantecWarmupPlan
	}
	if !IsNil(o.FortinetWarmupPlan) {
		toSerialize["fortinetWarmupPlan"] = o.FortinetWarmupPlan
	}
	if !IsNil(o.SophosWarmupPlan) {
		toSerialize["sophosWarmupPlan"] = o.SophosWarmupPlan
	}
	if !IsNil(o.TrendmicroWarmupPlan) {
		toSerialize["trendmicroWarmupPlan"] = o.TrendmicroWarmupPlan
	}
	if !IsNil(o.CheckpointWarmupPlan) {
		toSerialize["checkpointWarmupPlan"] = o.CheckpointWarmupPlan
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.WarmupInterval) {
		toSerialize["warmupInterval"] = o.WarmupInterval
	}
	return toSerialize, nil
}

type NullableAutoWarmupPlan struct {
	value *AutoWarmupPlan
	isSet bool
}

func (v NullableAutoWarmupPlan) Get() *AutoWarmupPlan {
	return v.value
}

func (v *NullableAutoWarmupPlan) Set(val *AutoWarmupPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoWarmupPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoWarmupPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoWarmupPlan(val *AutoWarmupPlan) *NullableAutoWarmupPlan {
	return &NullableAutoWarmupPlan{value: val, isSet: true}
}

func (v NullableAutoWarmupPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoWarmupPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


