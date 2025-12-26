/*
SendPost API

# Introduction  SendPost provides email API and SMTP relay which can be used not just to send & measure but also alert & optimised email sending.  You can use SendPost to:  * Send personalised emails to multiple recipients using email API   * Track opens and clicks  * Analyse statistics around open, clicks, bounce, unsubscribe and spam    At and advanced level you can use it to:  * Manage multiple sub-accounts which may map to your promotional or transactional sending, multiple product lines or multiple customers   * Classify your emails using groups for better analysis  * Analyse and fix email sending at sub-account level, IP Pool level or group level  * Have automated alerts to notify disruptions regarding email sending  * Manage different dedicated IP Pools so to better control your email sending  * Automatically know when IP or domain is blacklisted or sender score is down  * Leverage pro deliverability tools to get significantly better email deliverability & inboxing   [<img src=\"https://run.pstmn.io/button.svg\" alt=\"Run In Postman\" style=\"width: 128px; height: 32px;\">](https://god.gw.postman.com/run-collection/33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b%26entityType%3Dcollection%26workspaceId%3D6b1e4f65-96a9-4136-9512-6266c852517e)   # Overview  ## REST API  SendPost API is built on REST API principles. Authenticated users can interact with any of the API endpoints to perform:  * **GET**- to get a resource  * **POST** - to create a resource  * **PUT** - to update an existing resource  * **DELETE** - to delete a resource   The API endpoint for all API calls is: <code>https://api.sendpost.io/api/v1</code>   Some conventions that have been followed in the API design overall are following:   * All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header.  * All resource endpoints end with singular name and not plural. So we have <code>domain</code> instead of domains for domain resource endpoint. Likewise we have <code>sender</code> instead of senders for sender resource endpoint.  * Body submitted for POST / PUT API calls as well as JSON response from SendPost API follow camelcase convention  * All timestamps returned in response (created or submittedAt response fields) are UNIX nano epoch timestamp.   <aside class=\"success\"> All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header. </aside>   SendPost uses conventional HTTP response codes to indicate the success or failure of an API request.    * Codes in the <code>2xx</code> range indicate success.   * Codes in the <code>4xx</code> range indicate an error owing due to unauthorize access, incorrect request parameters or body etc.  * Code in the <code>5xx</code> range indicate an eror with SendPost's servers ( internal service issue or maintenance )   <aside class=\"info\"> SendPost all responses return <code>created</code> in UNIX nano epoch timestamp.  </aside>   ## Authentication  SendPost uses API keys for authentication. You can register a new SendPost API key at our [developer portal](https://app.sendpost.io/register).   SendPost expects the API key to be included in all API requests to the server in a header that looks like the following:   `X-SubAccount-ApiKey: AHEZEP8192SEGH`   This API key is used for all Sub-Account level operations such as:  * Sending emails  * Retrieving stats regarding open, click, bounce, unsubscribe and spam  * Uploading suppressions list  * Verifying sending domains and more  In addition to <code>X-SubAccount-ApiKey</code> you also have another API Key <code>X-Account-APIKey</code> which is used for Account level operations such as :  * Creating and managing sub-accounts  * Allocating IPs for your account  * Getting overall billing and usage information  * Email List validation  * Creating and managing alerts and more   <aside class=\"notice\"> You must look at individual API reference page to look at whether <code>X-SubAccount-ApiKey</code> is required or <code>X-Account-ApiKey</code> </aside>   In case an incorrect API Key header is specified or if it is missed you will get HTTP Response 401 ( Unauthorized ) response from SendPost.   ## HTTP Response Headers   Code           | Reason                 | Details ---------------| -----------------------| ----------- 200            | Success                | Everything went well 401            | Unauthorized           | Incorrect or missing API header either <code>X-SubAccount-ApiKey</code> or <code>X-Account-ApiKey</code> 403            | Forbidden              | Typically sent when resource with same name or details already exist 406            | Missing resource id    | Resource id specified is either missing or doesn't exist 422            | Unprocessable entity   | Request body is not in proper format 500            | Internal server error  | Some error happened at SendPost while processing API request 503            | Service Unavailable    | SendPost is offline for maintenance. Please try again later  # API SDKs  We have native SendPost SDKs in the following programming languages. You can integrate with them or create your own SDK with our API specification. In case you need any assistance with respect to API then do reachout to our team from website chat or email us at **hello@sendpost.io**   * [PHP](https://github.com/sendpost/sendpost_php_sdk)  * [Javascript](https://github.com/sendpost/sendpost_javascript_sdk)  * [Ruby](https://github.com/sendpost/sendpost_ruby_sdk)  * [Python](https://github.com/sendpost/sendpost_python_sdk)  * [Golang](https://github.com/sendpost/sendpost_go_sdk)   # API Reference  SendX REST API can be broken down into two major sub-sections:   * Sub-Account  * Account    Sub-Account API operations enable common email sending API use-cases like sending bulk email, adding new domains or senders for email sending programmatically, retrieving stats, adding suppressions etc. All Sub-Account API operations need to pass <code>X-SubAccount-ApiKey</code> header with every API call.   The Account API operations allow users to manage multiple sub-accounts and manage IPs. A single parent SendPost account can have 100's of sub-accounts. You may want to create sub-accounts for different products your company is running or to segregate types of emails or for managing email sending across multiple customers of yours.   # SMTP Reference  Simple Mail Transfer Protocol (SMTP) is a quick and easy way to send email from one server to another. SendPost provides an SMTP service that allows you to deliver your email via our servers instead of your own client or server.  This means you can count on SendPost's delivery at scale for your SMTP needs.    ## Integrating SMTP    1. Get the SMTP `username` and `password` from your SendPost account.  2. Set the server host in your email client or application to `smtp.sendpost.io`. This setting is sometimes referred to as the external SMTP server or the SMTP relay.  3. Set the `username` and `password`.  4. Set the port to `587` (or as specified below).  ## SMTP Ports   - For an unencrypted or a TLS connection, use port `25`, `2525` or `587`.  - For a SSL connection, use port `465`  - Check your firewall and network to ensure they're not blocking any of our SMTP Endpoints.   SendPost supports STARTTLS for establishing a TLS-encrypted connection. STARTTLS is a means of upgrading an unencrypted connection to an encrypted connection. There are versions of STARTTLS for a variety of protocols; the SMTP version is defined in [RFC 3207](https://www.ietf.org/rfc/rfc3207.txt).   To set up a STARTTLS connection, the SMTP client connects to the SendPost SMTP endpoint `smtp.sendpost.io` on port 25, 587, or 2525, issues an EHLO command, and waits for the server to announce that it supports the STARTTLS SMTP extension. The client then issues the STARTTLS command, initiating TLS negotiation. When negotiation is complete, the client issues an EHLO command over the new encrypted connection, and the SMTP session proceeds normally.   <aside class=\"success\"> If you are unsure which port to use, a TLS connection on port 587 is typically recommended. </aside>   ## Sending email from your application   ```javascript \"use strict\";  const nodemailer = require(\"nodemailer\");  async function main() { // create reusable transporter object using the default SMTP transport let transporter = nodemailer.createTransport({ host: \"smtp.sendpost.io\", port: 587, secure: false, // true for 465, false for other ports auth: { user:  \"<username>\" , // generated ethereal user pass: \"<password>\", // generated ethereal password }, requireTLS: true, debug: true, logger: true, });  // send mail with defined transport object try { let info = await transporter.sendMail({ from: 'erlich@piedpiper.com', to: 'gilfoyle@piedpiper.com', subject: 'Test Email Subject', html: '<h1>Hello Geeks!!!</h1>', }); console.log(\"Message sent: %s\", info.messageId); } catch (e) { console.log(e) } }  main().catch(console.error); ```  For PHP   ```php <?php // Import PHPMailer classes into the global namespace use PHPMailer\\PHPMailer\\PHPMailer; use PHPMailer\\PHPMailer\\SMTP; use PHPMailer\\PHPMailer\\Exception;  // Load Composer's autoloader require 'vendor/autoload.php';  $mail = new PHPMailer(true);  // Settings try { $mail->SMTPDebug = SMTP::DEBUG_CONNECTION;                  // Enable verbose debug output $mail->isSMTP();                                            // Send using SMTP $mail->Host       = 'smtp.sendpost.io';                     // Set the SMTP server to send through $mail->SMTPAuth   = true;                                   // Enable SMTP authentication $mail->Username   = '<username>';                           // SMTP username $mail->Password   = '<password>';                           // SMTP password $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;         // Enable implicit TLS encryption $mail->Port       = 587;                                    // TCP port to connect to; use 587 if you have set `SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS`  //Recipients $mail->setFrom('erlich@piedpiper.com', 'Erlich'); $mail->addAddress('gilfoyle@piedpiper.com', 'Gilfoyle');  //Content $mail->isHTML(true);                                  //Set email format to HTML $mail->Subject = 'Here is the subject'; $mail->Body    = 'This is the HTML message body <b>in bold!</b>'; $mail->AltBody = 'This is the body in plain text for non-HTML mail clients';  $mail->send(); echo 'Message has been sent';  } catch (Exception $e) { echo \"Message could not be sent. Mailer Error: {$mail->ErrorInfo}\"; } ``` For Python ```python #!/usr/bin/python3  import sys import os import re  from smtplib import SMTP import ssl  from email.mime.text import MIMEText  SMTPserver = 'smtp.sendpost.io' PORT = 587 sender =     'erlich@piedpiper.com' destination = ['gilfoyle@piedpiper.com']  USERNAME = \"<username>\" PASSWORD = \"<password>\"  # typical values for text_subtype are plain, html, xml text_subtype = 'plain'  content=\"\"\"\\ Test message \"\"\"  subject=\"Sent from Python\"  try: msg = MIMEText(content, text_subtype) msg['Subject']= subject msg['From']   = sender  conn = SMTP(SMTPserver, PORT) conn.ehlo() context = ssl.create_default_context() conn.starttls(context=context)  # upgrade to tls conn.ehlo() conn.set_debuglevel(True) conn.login(USERNAME, PASSWORD)  try: resp = conn.sendmail(sender, destination, msg.as_string()) print(\"Send Mail Response: \", resp) except Exception as e: print(\"Send Email Error: \", e) finally: conn.quit()  except Exception as e: print(\"Error:\", e) ``` For Golang ```go package main  import ( \"fmt\" \"net/smtp\" \"os\" )  // Sending Email Using Smtp in Golang  func main() {  username := \"<username>\" password := \"<password>\"  from := \"erlich@piedpiper.com\" toList := []string{\"gilfoyle@piedpiper.com\"} host := \"smtp.sendpost.io\" port := \"587\" // recommended  // This is the message to send in the mail msg := \"Hello geeks!!!\"  // We can't send strings directly in mail, // strings need to be converted into slice bytes body := []byte(msg)  // PlainAuth uses the given username and password to // authenticate to host and act as identity. // Usually identity should be the empty string, // to act as username. auth := smtp.PlainAuth(\"\", username, password, host)  // SendMail uses TLS connection to send the mail // The email is sent to all address in the toList, // the body should be of type bytes, not strings // This returns error if any occured. err := smtp.SendMail(host+\":\"+port, auth, from, toList, body)  // handling the errors if err != nil { fmt.Println(err) os.Exit(1) }  fmt.Println(\"Successfully sent mail to all user in toList\") }  ``` For Java ```java // implementation 'com.sun.mail:javax.mail:1.6.2'  import java.util.Properties;  import javax.mail.Message; import javax.mail.Session; import javax.mail.Transport; import javax.mail.internet.InternetAddress; import javax.mail.internet.MimeMessage;  public class SMTPConnect {  // This address must be verified. static final String FROM = \"erlich@piedpiper.com\"; static final String FROMNAME = \"Erlich Bachman\";  // Replace recipient@example.com with a \"To\" address. If your account // is still in the sandbox, this address must be verified. static final String TO = \"gilfoyle@piedpiper.com\";  // Replace smtp_username with your SendPost SMTP user name. static final String SMTP_USERNAME = \"<username>\";  // Replace smtp_password with your SendPost SMTP password. static final String SMTP_PASSWORD = \"<password>\";  // SMTP Host Name static final String HOST = \"smtp.sendpost.io\";  // The port you will connect to on SendPost SMTP Endpoint. static final int PORT = 587;  static final String SUBJECT = \"SendPost SMTP Test (SMTP interface accessed using Java)\";  static final String BODY = String.join( System.getProperty(\"line.separator\"), \"<h1>SendPost SMTP Test</h1>\", \"<p>This email was sent with SendPost using the \", \"<a href='https://github.com/eclipse-ee4j/mail'>Javamail Package</a>\", \" for <a href='https://www.java.com'>Java</a>.\" );  public static void main(String[] args) throws Exception {  // Create a Properties object to contain connection configuration information. Properties props = System.getProperties(); props.put(\"mail.transport.protocol\", \"smtp\"); props.put(\"mail.smtp.port\", PORT); props.put(\"mail.smtp.starttls.enable\", \"true\"); props.put(\"mail.smtp.debug\", \"true\"); props.put(\"mail.smtp.auth\", \"true\");  // Create a Session object to represent a mail session with the specified properties. Session session = Session.getDefaultInstance(props);  // Create a message with the specified information. MimeMessage msg = new MimeMessage(session); msg.setFrom(new InternetAddress(FROM,FROMNAME)); msg.setRecipient(Message.RecipientType.TO, new InternetAddress(TO)); msg.setSubject(SUBJECT); msg.setContent(BODY,\"text/html\");  // Create a transport. Transport transport = session.getTransport();  // Send the message. try { System.out.println(\"Sending...\");  // Connect to SendPost SMTP using the SMTP username and password you specified above. transport.connect(HOST, SMTP_USERNAME, SMTP_PASSWORD);  // Send the email. transport.sendMessage(msg, msg.getAllRecipients()); System.out.println(\"Email sent!\");  } catch (Exception ex) {  System.out.println(\"The email was not sent.\"); System.out.println(\"Error message: \" + ex.getMessage()); System.out.println(ex); } // Close and terminate the connection. } } ```  Many programming languages support sending email using SMTP. This capability might be built into the programming language itself, or it might be available as an add-on, plug-in, or library. You can take advantage of this capability by sending email through SendPost from within application programs that you write.  We have provided examples in Python3, Golang, Java, PHP, JS. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the IP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IP{}

// IP struct for IP
type IP struct {
	// Unique ID for the IP
	Id int32 `json:"id"`
	// The public IP address associated with the resource
	PublicIP string `json:"publicIP"`
	SystemDomain *Domain `json:"systemDomain,omitempty"`
	// The reverse DNS hostname for the IP
	ReverseDNSHostname *string `json:"reverseDNSHostname,omitempty"`
	// Type of the IP
	Type *int32 `json:"type,omitempty"`
	// Configuration for Gmail delivery settings in JSON format
	GmailSettings *string `json:"gmailSettings,omitempty"`
	// Configuration for Yahoo delivery settings in JSON format
	YahooSettings *string `json:"yahooSettings,omitempty"`
	// Configuration for AOL delivery settings in JSON format
	AolSettings *string `json:"aolSettings,omitempty"`
	// Configuration for Microsoft delivery settings in JSON format
	MicrosoftSettings *string `json:"microsoftSettings,omitempty"`
	// Configuration for Comcast delivery settings in JSON format
	ComcastSettings *string `json:"comcastSettings,omitempty"`
	// Configuration for Yandex delivery settings in JSON format
	YandexSettings *string `json:"yandexSettings,omitempty"`
	// Configuration for GMX delivery settings in JSON format
	GmxSettings *string `json:"gmxSettings,omitempty"`
	// Configuration for Mail.ru delivery settings in JSON format
	MailruSettings *string `json:"mailruSettings,omitempty"`
	// Configuration for iCloud delivery settings in JSON format
	IcloudSettings *string `json:"icloudSettings,omitempty"`
	// Configuration for Zoho delivery settings in JSON format
	ZohoSettings *string `json:"zohoSettings,omitempty"`
	// Configuration for QQ delivery settings in JSON format
	QqSettings *string `json:"qqSettings,omitempty"`
	// Default delivery settings in JSON format
	DefaultSettings *string `json:"defaultSettings,omitempty"`
	// Configuration for AT&T delivery settings in JSON format
	AttSettings *string `json:"attSettings,omitempty"`
	// Configuration for Office365 delivery settings in JSON format
	Office365Settings *string `json:"office365Settings,omitempty"`
	// Configuration for Google Workspace delivery settings in JSON format
	GoogleworkspaceSettings *string `json:"googleworkspaceSettings,omitempty"`
	// Configuration for Proofpoint delivery settings in JSON format
	ProofpointSettings *string `json:"proofpointSettings,omitempty"`
	// Configuration for Mimecast delivery settings in JSON format
	MimecastSettings *string `json:"mimecastSettings,omitempty"`
	// Configuration for Barracuda delivery settings in JSON format
	BarracudaSettings *string `json:"barracudaSettings,omitempty"`
	// Configuration for Cisco IronPort delivery settings in JSON format
	CiscoironportSettings *string `json:"ciscoironportSettings,omitempty"`
	// Configuration for Rackspace delivery settings in JSON format
	RackspaceSettings *string `json:"rackspaceSettings,omitempty"`
	// Configuration for Zoho Business delivery settings in JSON format
	ZohobusinessSettings *string `json:"zohobusinessSettings,omitempty"`
	// Configuration for Amazon WorkMail delivery settings in JSON format
	AmazonworkmailSettings *string `json:"amazonworkmailSettings,omitempty"`
	// Configuration for Symantec delivery settings in JSON format
	SymantecSettings *string `json:"symantecSettings,omitempty"`
	// Configuration for Fortinet delivery settings in JSON format
	FortinetSettings *string `json:"fortinetSettings,omitempty"`
	// Configuration for Sophos delivery settings in JSON format
	SophosSettings *string `json:"sophosSettings,omitempty"`
	// Configuration for TrendMicro delivery settings in JSON format
	TrendmicroSettings *string `json:"trendmicroSettings,omitempty"`
	// Configuration for CheckPoint delivery settings in JSON format
	CheckpointSettings *string `json:"checkpointSettings,omitempty"`
	// The timestamp (UNIX epoch) when the IP was created
	Created int64 `json:"created"`
	// Classification of the infrastructure
	InfraClassification *string `json:"infraClassification,omitempty"`
	// Indicates whether infrastructure monitoring is enabled
	InfraMonitor *bool `json:"infraMonitor,omitempty"`
	// The state of the IP
	State *int32 `json:"state,omitempty"`
	// The auto-warmup plan associated with the IP. Can be null if no warmup plan is assigned.
	AutoWarmupPlan *AutoWarmupPlan `json:"autoWarmupPlan,omitempty"`
	// Labels associated with the IP
	Labels []Label `json:"labels,omitempty"`
}

type _IP IP

// NewIP instantiates a new IP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIP(id int32, publicIP string, created int64) *IP {
	this := IP{}
	this.Id = id
	this.PublicIP = publicIP
	this.Created = created
	return &this
}

// NewIPWithDefaults instantiates a new IP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPWithDefaults() *IP {
	this := IP{}
	return &this
}

// GetId returns the Id field value
func (o *IP) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IP) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IP) SetId(v int32) {
	o.Id = v
}

// GetPublicIP returns the PublicIP field value
func (o *IP) GetPublicIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicIP
}

// GetPublicIPOk returns a tuple with the PublicIP field value
// and a boolean to check if the value has been set.
func (o *IP) GetPublicIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicIP, true
}

// SetPublicIP sets field value
func (o *IP) SetPublicIP(v string) {
	o.PublicIP = v
}

// GetSystemDomain returns the SystemDomain field value if set, zero value otherwise.
func (o *IP) GetSystemDomain() Domain {
	if o == nil || IsNil(o.SystemDomain) {
		var ret Domain
		return ret
	}
	return *o.SystemDomain
}

// GetSystemDomainOk returns a tuple with the SystemDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetSystemDomainOk() (*Domain, bool) {
	if o == nil || IsNil(o.SystemDomain) {
		return nil, false
	}
	return o.SystemDomain, true
}

// HasSystemDomain returns a boolean if a field has been set.
func (o *IP) HasSystemDomain() bool {
	if o != nil && !IsNil(o.SystemDomain) {
		return true
	}

	return false
}

// SetSystemDomain gets a reference to the given Domain and assigns it to the SystemDomain field.
func (o *IP) SetSystemDomain(v Domain) {
	o.SystemDomain = &v
}

// GetReverseDNSHostname returns the ReverseDNSHostname field value if set, zero value otherwise.
func (o *IP) GetReverseDNSHostname() string {
	if o == nil || IsNil(o.ReverseDNSHostname) {
		var ret string
		return ret
	}
	return *o.ReverseDNSHostname
}

// GetReverseDNSHostnameOk returns a tuple with the ReverseDNSHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetReverseDNSHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ReverseDNSHostname) {
		return nil, false
	}
	return o.ReverseDNSHostname, true
}

// HasReverseDNSHostname returns a boolean if a field has been set.
func (o *IP) HasReverseDNSHostname() bool {
	if o != nil && !IsNil(o.ReverseDNSHostname) {
		return true
	}

	return false
}

// SetReverseDNSHostname gets a reference to the given string and assigns it to the ReverseDNSHostname field.
func (o *IP) SetReverseDNSHostname(v string) {
	o.ReverseDNSHostname = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IP) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IP) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *IP) SetType(v int32) {
	o.Type = &v
}

// GetGmailSettings returns the GmailSettings field value if set, zero value otherwise.
func (o *IP) GetGmailSettings() string {
	if o == nil || IsNil(o.GmailSettings) {
		var ret string
		return ret
	}
	return *o.GmailSettings
}

// GetGmailSettingsOk returns a tuple with the GmailSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetGmailSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.GmailSettings) {
		return nil, false
	}
	return o.GmailSettings, true
}

// HasGmailSettings returns a boolean if a field has been set.
func (o *IP) HasGmailSettings() bool {
	if o != nil && !IsNil(o.GmailSettings) {
		return true
	}

	return false
}

// SetGmailSettings gets a reference to the given string and assigns it to the GmailSettings field.
func (o *IP) SetGmailSettings(v string) {
	o.GmailSettings = &v
}

// GetYahooSettings returns the YahooSettings field value if set, zero value otherwise.
func (o *IP) GetYahooSettings() string {
	if o == nil || IsNil(o.YahooSettings) {
		var ret string
		return ret
	}
	return *o.YahooSettings
}

// GetYahooSettingsOk returns a tuple with the YahooSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetYahooSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.YahooSettings) {
		return nil, false
	}
	return o.YahooSettings, true
}

// HasYahooSettings returns a boolean if a field has been set.
func (o *IP) HasYahooSettings() bool {
	if o != nil && !IsNil(o.YahooSettings) {
		return true
	}

	return false
}

// SetYahooSettings gets a reference to the given string and assigns it to the YahooSettings field.
func (o *IP) SetYahooSettings(v string) {
	o.YahooSettings = &v
}

// GetAolSettings returns the AolSettings field value if set, zero value otherwise.
func (o *IP) GetAolSettings() string {
	if o == nil || IsNil(o.AolSettings) {
		var ret string
		return ret
	}
	return *o.AolSettings
}

// GetAolSettingsOk returns a tuple with the AolSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetAolSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.AolSettings) {
		return nil, false
	}
	return o.AolSettings, true
}

// HasAolSettings returns a boolean if a field has been set.
func (o *IP) HasAolSettings() bool {
	if o != nil && !IsNil(o.AolSettings) {
		return true
	}

	return false
}

// SetAolSettings gets a reference to the given string and assigns it to the AolSettings field.
func (o *IP) SetAolSettings(v string) {
	o.AolSettings = &v
}

// GetMicrosoftSettings returns the MicrosoftSettings field value if set, zero value otherwise.
func (o *IP) GetMicrosoftSettings() string {
	if o == nil || IsNil(o.MicrosoftSettings) {
		var ret string
		return ret
	}
	return *o.MicrosoftSettings
}

// GetMicrosoftSettingsOk returns a tuple with the MicrosoftSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetMicrosoftSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.MicrosoftSettings) {
		return nil, false
	}
	return o.MicrosoftSettings, true
}

// HasMicrosoftSettings returns a boolean if a field has been set.
func (o *IP) HasMicrosoftSettings() bool {
	if o != nil && !IsNil(o.MicrosoftSettings) {
		return true
	}

	return false
}

// SetMicrosoftSettings gets a reference to the given string and assigns it to the MicrosoftSettings field.
func (o *IP) SetMicrosoftSettings(v string) {
	o.MicrosoftSettings = &v
}

// GetComcastSettings returns the ComcastSettings field value if set, zero value otherwise.
func (o *IP) GetComcastSettings() string {
	if o == nil || IsNil(o.ComcastSettings) {
		var ret string
		return ret
	}
	return *o.ComcastSettings
}

// GetComcastSettingsOk returns a tuple with the ComcastSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetComcastSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.ComcastSettings) {
		return nil, false
	}
	return o.ComcastSettings, true
}

// HasComcastSettings returns a boolean if a field has been set.
func (o *IP) HasComcastSettings() bool {
	if o != nil && !IsNil(o.ComcastSettings) {
		return true
	}

	return false
}

// SetComcastSettings gets a reference to the given string and assigns it to the ComcastSettings field.
func (o *IP) SetComcastSettings(v string) {
	o.ComcastSettings = &v
}

// GetYandexSettings returns the YandexSettings field value if set, zero value otherwise.
func (o *IP) GetYandexSettings() string {
	if o == nil || IsNil(o.YandexSettings) {
		var ret string
		return ret
	}
	return *o.YandexSettings
}

// GetYandexSettingsOk returns a tuple with the YandexSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetYandexSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.YandexSettings) {
		return nil, false
	}
	return o.YandexSettings, true
}

// HasYandexSettings returns a boolean if a field has been set.
func (o *IP) HasYandexSettings() bool {
	if o != nil && !IsNil(o.YandexSettings) {
		return true
	}

	return false
}

// SetYandexSettings gets a reference to the given string and assigns it to the YandexSettings field.
func (o *IP) SetYandexSettings(v string) {
	o.YandexSettings = &v
}

// GetGmxSettings returns the GmxSettings field value if set, zero value otherwise.
func (o *IP) GetGmxSettings() string {
	if o == nil || IsNil(o.GmxSettings) {
		var ret string
		return ret
	}
	return *o.GmxSettings
}

// GetGmxSettingsOk returns a tuple with the GmxSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetGmxSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.GmxSettings) {
		return nil, false
	}
	return o.GmxSettings, true
}

// HasGmxSettings returns a boolean if a field has been set.
func (o *IP) HasGmxSettings() bool {
	if o != nil && !IsNil(o.GmxSettings) {
		return true
	}

	return false
}

// SetGmxSettings gets a reference to the given string and assigns it to the GmxSettings field.
func (o *IP) SetGmxSettings(v string) {
	o.GmxSettings = &v
}

// GetMailruSettings returns the MailruSettings field value if set, zero value otherwise.
func (o *IP) GetMailruSettings() string {
	if o == nil || IsNil(o.MailruSettings) {
		var ret string
		return ret
	}
	return *o.MailruSettings
}

// GetMailruSettingsOk returns a tuple with the MailruSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetMailruSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.MailruSettings) {
		return nil, false
	}
	return o.MailruSettings, true
}

// HasMailruSettings returns a boolean if a field has been set.
func (o *IP) HasMailruSettings() bool {
	if o != nil && !IsNil(o.MailruSettings) {
		return true
	}

	return false
}

// SetMailruSettings gets a reference to the given string and assigns it to the MailruSettings field.
func (o *IP) SetMailruSettings(v string) {
	o.MailruSettings = &v
}

// GetIcloudSettings returns the IcloudSettings field value if set, zero value otherwise.
func (o *IP) GetIcloudSettings() string {
	if o == nil || IsNil(o.IcloudSettings) {
		var ret string
		return ret
	}
	return *o.IcloudSettings
}

// GetIcloudSettingsOk returns a tuple with the IcloudSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetIcloudSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.IcloudSettings) {
		return nil, false
	}
	return o.IcloudSettings, true
}

// HasIcloudSettings returns a boolean if a field has been set.
func (o *IP) HasIcloudSettings() bool {
	if o != nil && !IsNil(o.IcloudSettings) {
		return true
	}

	return false
}

// SetIcloudSettings gets a reference to the given string and assigns it to the IcloudSettings field.
func (o *IP) SetIcloudSettings(v string) {
	o.IcloudSettings = &v
}

// GetZohoSettings returns the ZohoSettings field value if set, zero value otherwise.
func (o *IP) GetZohoSettings() string {
	if o == nil || IsNil(o.ZohoSettings) {
		var ret string
		return ret
	}
	return *o.ZohoSettings
}

// GetZohoSettingsOk returns a tuple with the ZohoSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetZohoSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.ZohoSettings) {
		return nil, false
	}
	return o.ZohoSettings, true
}

// HasZohoSettings returns a boolean if a field has been set.
func (o *IP) HasZohoSettings() bool {
	if o != nil && !IsNil(o.ZohoSettings) {
		return true
	}

	return false
}

// SetZohoSettings gets a reference to the given string and assigns it to the ZohoSettings field.
func (o *IP) SetZohoSettings(v string) {
	o.ZohoSettings = &v
}

// GetQqSettings returns the QqSettings field value if set, zero value otherwise.
func (o *IP) GetQqSettings() string {
	if o == nil || IsNil(o.QqSettings) {
		var ret string
		return ret
	}
	return *o.QqSettings
}

// GetQqSettingsOk returns a tuple with the QqSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetQqSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.QqSettings) {
		return nil, false
	}
	return o.QqSettings, true
}

// HasQqSettings returns a boolean if a field has been set.
func (o *IP) HasQqSettings() bool {
	if o != nil && !IsNil(o.QqSettings) {
		return true
	}

	return false
}

// SetQqSettings gets a reference to the given string and assigns it to the QqSettings field.
func (o *IP) SetQqSettings(v string) {
	o.QqSettings = &v
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *IP) GetDefaultSettings() string {
	if o == nil || IsNil(o.DefaultSettings) {
		var ret string
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetDefaultSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultSettings) {
		return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *IP) HasDefaultSettings() bool {
	if o != nil && !IsNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given string and assigns it to the DefaultSettings field.
func (o *IP) SetDefaultSettings(v string) {
	o.DefaultSettings = &v
}

// GetAttSettings returns the AttSettings field value if set, zero value otherwise.
func (o *IP) GetAttSettings() string {
	if o == nil || IsNil(o.AttSettings) {
		var ret string
		return ret
	}
	return *o.AttSettings
}

// GetAttSettingsOk returns a tuple with the AttSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetAttSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.AttSettings) {
		return nil, false
	}
	return o.AttSettings, true
}

// HasAttSettings returns a boolean if a field has been set.
func (o *IP) HasAttSettings() bool {
	if o != nil && !IsNil(o.AttSettings) {
		return true
	}

	return false
}

// SetAttSettings gets a reference to the given string and assigns it to the AttSettings field.
func (o *IP) SetAttSettings(v string) {
	o.AttSettings = &v
}

// GetOffice365Settings returns the Office365Settings field value if set, zero value otherwise.
func (o *IP) GetOffice365Settings() string {
	if o == nil || IsNil(o.Office365Settings) {
		var ret string
		return ret
	}
	return *o.Office365Settings
}

// GetOffice365SettingsOk returns a tuple with the Office365Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetOffice365SettingsOk() (*string, bool) {
	if o == nil || IsNil(o.Office365Settings) {
		return nil, false
	}
	return o.Office365Settings, true
}

// HasOffice365Settings returns a boolean if a field has been set.
func (o *IP) HasOffice365Settings() bool {
	if o != nil && !IsNil(o.Office365Settings) {
		return true
	}

	return false
}

// SetOffice365Settings gets a reference to the given string and assigns it to the Office365Settings field.
func (o *IP) SetOffice365Settings(v string) {
	o.Office365Settings = &v
}

// GetGoogleworkspaceSettings returns the GoogleworkspaceSettings field value if set, zero value otherwise.
func (o *IP) GetGoogleworkspaceSettings() string {
	if o == nil || IsNil(o.GoogleworkspaceSettings) {
		var ret string
		return ret
	}
	return *o.GoogleworkspaceSettings
}

// GetGoogleworkspaceSettingsOk returns a tuple with the GoogleworkspaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetGoogleworkspaceSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleworkspaceSettings) {
		return nil, false
	}
	return o.GoogleworkspaceSettings, true
}

// HasGoogleworkspaceSettings returns a boolean if a field has been set.
func (o *IP) HasGoogleworkspaceSettings() bool {
	if o != nil && !IsNil(o.GoogleworkspaceSettings) {
		return true
	}

	return false
}

// SetGoogleworkspaceSettings gets a reference to the given string and assigns it to the GoogleworkspaceSettings field.
func (o *IP) SetGoogleworkspaceSettings(v string) {
	o.GoogleworkspaceSettings = &v
}

// GetProofpointSettings returns the ProofpointSettings field value if set, zero value otherwise.
func (o *IP) GetProofpointSettings() string {
	if o == nil || IsNil(o.ProofpointSettings) {
		var ret string
		return ret
	}
	return *o.ProofpointSettings
}

// GetProofpointSettingsOk returns a tuple with the ProofpointSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetProofpointSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.ProofpointSettings) {
		return nil, false
	}
	return o.ProofpointSettings, true
}

// HasProofpointSettings returns a boolean if a field has been set.
func (o *IP) HasProofpointSettings() bool {
	if o != nil && !IsNil(o.ProofpointSettings) {
		return true
	}

	return false
}

// SetProofpointSettings gets a reference to the given string and assigns it to the ProofpointSettings field.
func (o *IP) SetProofpointSettings(v string) {
	o.ProofpointSettings = &v
}

// GetMimecastSettings returns the MimecastSettings field value if set, zero value otherwise.
func (o *IP) GetMimecastSettings() string {
	if o == nil || IsNil(o.MimecastSettings) {
		var ret string
		return ret
	}
	return *o.MimecastSettings
}

// GetMimecastSettingsOk returns a tuple with the MimecastSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetMimecastSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.MimecastSettings) {
		return nil, false
	}
	return o.MimecastSettings, true
}

// HasMimecastSettings returns a boolean if a field has been set.
func (o *IP) HasMimecastSettings() bool {
	if o != nil && !IsNil(o.MimecastSettings) {
		return true
	}

	return false
}

// SetMimecastSettings gets a reference to the given string and assigns it to the MimecastSettings field.
func (o *IP) SetMimecastSettings(v string) {
	o.MimecastSettings = &v
}

// GetBarracudaSettings returns the BarracudaSettings field value if set, zero value otherwise.
func (o *IP) GetBarracudaSettings() string {
	if o == nil || IsNil(o.BarracudaSettings) {
		var ret string
		return ret
	}
	return *o.BarracudaSettings
}

// GetBarracudaSettingsOk returns a tuple with the BarracudaSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetBarracudaSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.BarracudaSettings) {
		return nil, false
	}
	return o.BarracudaSettings, true
}

// HasBarracudaSettings returns a boolean if a field has been set.
func (o *IP) HasBarracudaSettings() bool {
	if o != nil && !IsNil(o.BarracudaSettings) {
		return true
	}

	return false
}

// SetBarracudaSettings gets a reference to the given string and assigns it to the BarracudaSettings field.
func (o *IP) SetBarracudaSettings(v string) {
	o.BarracudaSettings = &v
}

// GetCiscoironportSettings returns the CiscoironportSettings field value if set, zero value otherwise.
func (o *IP) GetCiscoironportSettings() string {
	if o == nil || IsNil(o.CiscoironportSettings) {
		var ret string
		return ret
	}
	return *o.CiscoironportSettings
}

// GetCiscoironportSettingsOk returns a tuple with the CiscoironportSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetCiscoironportSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.CiscoironportSettings) {
		return nil, false
	}
	return o.CiscoironportSettings, true
}

// HasCiscoironportSettings returns a boolean if a field has been set.
func (o *IP) HasCiscoironportSettings() bool {
	if o != nil && !IsNil(o.CiscoironportSettings) {
		return true
	}

	return false
}

// SetCiscoironportSettings gets a reference to the given string and assigns it to the CiscoironportSettings field.
func (o *IP) SetCiscoironportSettings(v string) {
	o.CiscoironportSettings = &v
}

// GetRackspaceSettings returns the RackspaceSettings field value if set, zero value otherwise.
func (o *IP) GetRackspaceSettings() string {
	if o == nil || IsNil(o.RackspaceSettings) {
		var ret string
		return ret
	}
	return *o.RackspaceSettings
}

// GetRackspaceSettingsOk returns a tuple with the RackspaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetRackspaceSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.RackspaceSettings) {
		return nil, false
	}
	return o.RackspaceSettings, true
}

// HasRackspaceSettings returns a boolean if a field has been set.
func (o *IP) HasRackspaceSettings() bool {
	if o != nil && !IsNil(o.RackspaceSettings) {
		return true
	}

	return false
}

// SetRackspaceSettings gets a reference to the given string and assigns it to the RackspaceSettings field.
func (o *IP) SetRackspaceSettings(v string) {
	o.RackspaceSettings = &v
}

// GetZohobusinessSettings returns the ZohobusinessSettings field value if set, zero value otherwise.
func (o *IP) GetZohobusinessSettings() string {
	if o == nil || IsNil(o.ZohobusinessSettings) {
		var ret string
		return ret
	}
	return *o.ZohobusinessSettings
}

// GetZohobusinessSettingsOk returns a tuple with the ZohobusinessSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetZohobusinessSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.ZohobusinessSettings) {
		return nil, false
	}
	return o.ZohobusinessSettings, true
}

// HasZohobusinessSettings returns a boolean if a field has been set.
func (o *IP) HasZohobusinessSettings() bool {
	if o != nil && !IsNil(o.ZohobusinessSettings) {
		return true
	}

	return false
}

// SetZohobusinessSettings gets a reference to the given string and assigns it to the ZohobusinessSettings field.
func (o *IP) SetZohobusinessSettings(v string) {
	o.ZohobusinessSettings = &v
}

// GetAmazonworkmailSettings returns the AmazonworkmailSettings field value if set, zero value otherwise.
func (o *IP) GetAmazonworkmailSettings() string {
	if o == nil || IsNil(o.AmazonworkmailSettings) {
		var ret string
		return ret
	}
	return *o.AmazonworkmailSettings
}

// GetAmazonworkmailSettingsOk returns a tuple with the AmazonworkmailSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetAmazonworkmailSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonworkmailSettings) {
		return nil, false
	}
	return o.AmazonworkmailSettings, true
}

// HasAmazonworkmailSettings returns a boolean if a field has been set.
func (o *IP) HasAmazonworkmailSettings() bool {
	if o != nil && !IsNil(o.AmazonworkmailSettings) {
		return true
	}

	return false
}

// SetAmazonworkmailSettings gets a reference to the given string and assigns it to the AmazonworkmailSettings field.
func (o *IP) SetAmazonworkmailSettings(v string) {
	o.AmazonworkmailSettings = &v
}

// GetSymantecSettings returns the SymantecSettings field value if set, zero value otherwise.
func (o *IP) GetSymantecSettings() string {
	if o == nil || IsNil(o.SymantecSettings) {
		var ret string
		return ret
	}
	return *o.SymantecSettings
}

// GetSymantecSettingsOk returns a tuple with the SymantecSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetSymantecSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.SymantecSettings) {
		return nil, false
	}
	return o.SymantecSettings, true
}

// HasSymantecSettings returns a boolean if a field has been set.
func (o *IP) HasSymantecSettings() bool {
	if o != nil && !IsNil(o.SymantecSettings) {
		return true
	}

	return false
}

// SetSymantecSettings gets a reference to the given string and assigns it to the SymantecSettings field.
func (o *IP) SetSymantecSettings(v string) {
	o.SymantecSettings = &v
}

// GetFortinetSettings returns the FortinetSettings field value if set, zero value otherwise.
func (o *IP) GetFortinetSettings() string {
	if o == nil || IsNil(o.FortinetSettings) {
		var ret string
		return ret
	}
	return *o.FortinetSettings
}

// GetFortinetSettingsOk returns a tuple with the FortinetSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetFortinetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.FortinetSettings) {
		return nil, false
	}
	return o.FortinetSettings, true
}

// HasFortinetSettings returns a boolean if a field has been set.
func (o *IP) HasFortinetSettings() bool {
	if o != nil && !IsNil(o.FortinetSettings) {
		return true
	}

	return false
}

// SetFortinetSettings gets a reference to the given string and assigns it to the FortinetSettings field.
func (o *IP) SetFortinetSettings(v string) {
	o.FortinetSettings = &v
}

// GetSophosSettings returns the SophosSettings field value if set, zero value otherwise.
func (o *IP) GetSophosSettings() string {
	if o == nil || IsNil(o.SophosSettings) {
		var ret string
		return ret
	}
	return *o.SophosSettings
}

// GetSophosSettingsOk returns a tuple with the SophosSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetSophosSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.SophosSettings) {
		return nil, false
	}
	return o.SophosSettings, true
}

// HasSophosSettings returns a boolean if a field has been set.
func (o *IP) HasSophosSettings() bool {
	if o != nil && !IsNil(o.SophosSettings) {
		return true
	}

	return false
}

// SetSophosSettings gets a reference to the given string and assigns it to the SophosSettings field.
func (o *IP) SetSophosSettings(v string) {
	o.SophosSettings = &v
}

// GetTrendmicroSettings returns the TrendmicroSettings field value if set, zero value otherwise.
func (o *IP) GetTrendmicroSettings() string {
	if o == nil || IsNil(o.TrendmicroSettings) {
		var ret string
		return ret
	}
	return *o.TrendmicroSettings
}

// GetTrendmicroSettingsOk returns a tuple with the TrendmicroSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetTrendmicroSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.TrendmicroSettings) {
		return nil, false
	}
	return o.TrendmicroSettings, true
}

// HasTrendmicroSettings returns a boolean if a field has been set.
func (o *IP) HasTrendmicroSettings() bool {
	if o != nil && !IsNil(o.TrendmicroSettings) {
		return true
	}

	return false
}

// SetTrendmicroSettings gets a reference to the given string and assigns it to the TrendmicroSettings field.
func (o *IP) SetTrendmicroSettings(v string) {
	o.TrendmicroSettings = &v
}

// GetCheckpointSettings returns the CheckpointSettings field value if set, zero value otherwise.
func (o *IP) GetCheckpointSettings() string {
	if o == nil || IsNil(o.CheckpointSettings) {
		var ret string
		return ret
	}
	return *o.CheckpointSettings
}

// GetCheckpointSettingsOk returns a tuple with the CheckpointSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetCheckpointSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.CheckpointSettings) {
		return nil, false
	}
	return o.CheckpointSettings, true
}

// HasCheckpointSettings returns a boolean if a field has been set.
func (o *IP) HasCheckpointSettings() bool {
	if o != nil && !IsNil(o.CheckpointSettings) {
		return true
	}

	return false
}

// SetCheckpointSettings gets a reference to the given string and assigns it to the CheckpointSettings field.
func (o *IP) SetCheckpointSettings(v string) {
	o.CheckpointSettings = &v
}

// GetCreated returns the Created field value
func (o *IP) GetCreated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IP) GetCreatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *IP) SetCreated(v int64) {
	o.Created = v
}

// GetInfraClassification returns the InfraClassification field value if set, zero value otherwise.
func (o *IP) GetInfraClassification() string {
	if o == nil || IsNil(o.InfraClassification) {
		var ret string
		return ret
	}
	return *o.InfraClassification
}

// GetInfraClassificationOk returns a tuple with the InfraClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetInfraClassificationOk() (*string, bool) {
	if o == nil || IsNil(o.InfraClassification) {
		return nil, false
	}
	return o.InfraClassification, true
}

// HasInfraClassification returns a boolean if a field has been set.
func (o *IP) HasInfraClassification() bool {
	if o != nil && !IsNil(o.InfraClassification) {
		return true
	}

	return false
}

// SetInfraClassification gets a reference to the given string and assigns it to the InfraClassification field.
func (o *IP) SetInfraClassification(v string) {
	o.InfraClassification = &v
}

// GetInfraMonitor returns the InfraMonitor field value if set, zero value otherwise.
func (o *IP) GetInfraMonitor() bool {
	if o == nil || IsNil(o.InfraMonitor) {
		var ret bool
		return ret
	}
	return *o.InfraMonitor
}

// GetInfraMonitorOk returns a tuple with the InfraMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetInfraMonitorOk() (*bool, bool) {
	if o == nil || IsNil(o.InfraMonitor) {
		return nil, false
	}
	return o.InfraMonitor, true
}

// HasInfraMonitor returns a boolean if a field has been set.
func (o *IP) HasInfraMonitor() bool {
	if o != nil && !IsNil(o.InfraMonitor) {
		return true
	}

	return false
}

// SetInfraMonitor gets a reference to the given bool and assigns it to the InfraMonitor field.
func (o *IP) SetInfraMonitor(v bool) {
	o.InfraMonitor = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IP) GetState() int32 {
	if o == nil || IsNil(o.State) {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetStateOk() (*int32, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IP) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *IP) SetState(v int32) {
	o.State = &v
}

// GetAutoWarmupPlan returns the AutoWarmupPlan field value if set, zero value otherwise.
func (o *IP) GetAutoWarmupPlan() AutoWarmupPlan {
	if o == nil || IsNil(o.AutoWarmupPlan) {
		var ret AutoWarmupPlan
		return ret
	}
	return *o.AutoWarmupPlan
}

// GetAutoWarmupPlanOk returns a tuple with the AutoWarmupPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetAutoWarmupPlanOk() (*AutoWarmupPlan, bool) {
	if o == nil || IsNil(o.AutoWarmupPlan) {
		return nil, false
	}
	return o.AutoWarmupPlan, true
}

// HasAutoWarmupPlan returns a boolean if a field has been set.
func (o *IP) HasAutoWarmupPlan() bool {
	if o != nil && !IsNil(o.AutoWarmupPlan) {
		return true
	}

	return false
}

// SetAutoWarmupPlan gets a reference to the given AutoWarmupPlan and assigns it to the AutoWarmupPlan field.
func (o *IP) SetAutoWarmupPlan(v AutoWarmupPlan) {
	o.AutoWarmupPlan = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *IP) GetLabels() []Label {
	if o == nil || IsNil(o.Labels) {
		var ret []Label
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IP) GetLabelsOk() ([]Label, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *IP) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *IP) SetLabels(v []Label) {
	o.Labels = v
}

func (o IP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["publicIP"] = o.PublicIP
	if !IsNil(o.SystemDomain) {
		toSerialize["systemDomain"] = o.SystemDomain
	}
	if !IsNil(o.ReverseDNSHostname) {
		toSerialize["reverseDNSHostname"] = o.ReverseDNSHostname
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.GmailSettings) {
		toSerialize["gmailSettings"] = o.GmailSettings
	}
	if !IsNil(o.YahooSettings) {
		toSerialize["yahooSettings"] = o.YahooSettings
	}
	if !IsNil(o.AolSettings) {
		toSerialize["aolSettings"] = o.AolSettings
	}
	if !IsNil(o.MicrosoftSettings) {
		toSerialize["microsoftSettings"] = o.MicrosoftSettings
	}
	if !IsNil(o.ComcastSettings) {
		toSerialize["comcastSettings"] = o.ComcastSettings
	}
	if !IsNil(o.YandexSettings) {
		toSerialize["yandexSettings"] = o.YandexSettings
	}
	if !IsNil(o.GmxSettings) {
		toSerialize["gmxSettings"] = o.GmxSettings
	}
	if !IsNil(o.MailruSettings) {
		toSerialize["mailruSettings"] = o.MailruSettings
	}
	if !IsNil(o.IcloudSettings) {
		toSerialize["icloudSettings"] = o.IcloudSettings
	}
	if !IsNil(o.ZohoSettings) {
		toSerialize["zohoSettings"] = o.ZohoSettings
	}
	if !IsNil(o.QqSettings) {
		toSerialize["qqSettings"] = o.QqSettings
	}
	if !IsNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !IsNil(o.AttSettings) {
		toSerialize["attSettings"] = o.AttSettings
	}
	if !IsNil(o.Office365Settings) {
		toSerialize["office365Settings"] = o.Office365Settings
	}
	if !IsNil(o.GoogleworkspaceSettings) {
		toSerialize["googleworkspaceSettings"] = o.GoogleworkspaceSettings
	}
	if !IsNil(o.ProofpointSettings) {
		toSerialize["proofpointSettings"] = o.ProofpointSettings
	}
	if !IsNil(o.MimecastSettings) {
		toSerialize["mimecastSettings"] = o.MimecastSettings
	}
	if !IsNil(o.BarracudaSettings) {
		toSerialize["barracudaSettings"] = o.BarracudaSettings
	}
	if !IsNil(o.CiscoironportSettings) {
		toSerialize["ciscoironportSettings"] = o.CiscoironportSettings
	}
	if !IsNil(o.RackspaceSettings) {
		toSerialize["rackspaceSettings"] = o.RackspaceSettings
	}
	if !IsNil(o.ZohobusinessSettings) {
		toSerialize["zohobusinessSettings"] = o.ZohobusinessSettings
	}
	if !IsNil(o.AmazonworkmailSettings) {
		toSerialize["amazonworkmailSettings"] = o.AmazonworkmailSettings
	}
	if !IsNil(o.SymantecSettings) {
		toSerialize["symantecSettings"] = o.SymantecSettings
	}
	if !IsNil(o.FortinetSettings) {
		toSerialize["fortinetSettings"] = o.FortinetSettings
	}
	if !IsNil(o.SophosSettings) {
		toSerialize["sophosSettings"] = o.SophosSettings
	}
	if !IsNil(o.TrendmicroSettings) {
		toSerialize["trendmicroSettings"] = o.TrendmicroSettings
	}
	if !IsNil(o.CheckpointSettings) {
		toSerialize["checkpointSettings"] = o.CheckpointSettings
	}
	toSerialize["created"] = o.Created
	if !IsNil(o.InfraClassification) {
		toSerialize["infraClassification"] = o.InfraClassification
	}
	if !IsNil(o.InfraMonitor) {
		toSerialize["infraMonitor"] = o.InfraMonitor
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.AutoWarmupPlan) {
		toSerialize["autoWarmupPlan"] = o.AutoWarmupPlan
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

func (o *IP) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"publicIP",
		"created",
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

	varIP := _IP{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIP)

	if err != nil {
		return err
	}

	*o = IP(varIP)

	return err
}

type NullableIP struct {
	value *IP
	isSet bool
}

func (v NullableIP) Get() *IP {
	return v.value
}

func (v *NullableIP) Set(val *IP) {
	v.value = val
	v.isSet = true
}

func (v NullableIP) IsSet() bool {
	return v.isSet
}

func (v *NullableIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIP(val *IP) *NullableIP {
	return &NullableIP{value: val, isSet: true}
}

func (v NullableIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


