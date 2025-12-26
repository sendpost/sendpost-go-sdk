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

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	// Unique ID for the domain.
	Id *int32 `json:"id,omitempty"`
	// Name of the domain.
	Name *string `json:"name,omitempty"`
	Dkim *DomainDkim `json:"dkim,omitempty"`
	Spf *DomainSpf `json:"spf,omitempty"`
	ReturnPath *DomainReturnPath `json:"returnPath,omitempty"`
	Track *DomainTrack `json:"track,omitempty"`
	Dmarc *DomainDmarc `json:"dmarc,omitempty"`
	// DKIM configuration
	DkimConfig *string `json:"dkimConfig,omitempty"`
	// Status of DKIM verification ( true or false )
	DkimVerified *bool `json:"dkimVerified,omitempty"`
	// Status of SPF verification ( true or false )
	SpfVerified *bool `json:"spfVerified,omitempty"`
	// Status of Mailbox verification ( true or false )
	MailboxVerified *bool `json:"mailboxVerified,omitempty"`
	// Status of DMARC verification ( true or false)
	DmarcVerified *bool `json:"dmarcVerified,omitempty"`
	// Status of ReturnPath verification ( true or false )
	ReturnPathVerified *bool `json:"returnPathVerified,omitempty"`
	// Status of Track verification ( true or false )
	TrackVerified *bool `json:"trackVerified,omitempty"`
	// Overall verification status of the domain
	Verified *bool `json:"verified,omitempty"`
	// Date when the domain was registered
	DomainRegisteredDate *string `json:"domainRegisteredDate,omitempty"`
	// UNIX epoch timestamp in nanoseconds.
	Created *int64 `json:"created,omitempty"`
	// Status of GPT verification ( true or false )
	GptVerified *bool `json:"gptVerified,omitempty"`
	Gpt *DomainGpt `json:"gpt,omitempty"`
	// Reason for DMARC verification failure
	DmarcFailureReason *string `json:"dmarcFailureReason,omitempty"`
	// Reason for DKIM verification failure
	DkimFailureReason *string `json:"dkimFailureReason,omitempty"`
	// Reason for Track verification failure
	TrackFailureReason *string `json:"trackFailureReason,omitempty"`
	// Reason for ReturnPath verification failure
	ReturnPathFailureReason *string `json:"returnPathFailureReason,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Domain) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Domain) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Domain) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Domain) SetName(v string) {
	o.Name = &v
}

// GetDkim returns the Dkim field value if set, zero value otherwise.
func (o *Domain) GetDkim() DomainDkim {
	if o == nil || IsNil(o.Dkim) {
		var ret DomainDkim
		return ret
	}
	return *o.Dkim
}

// GetDkimOk returns a tuple with the Dkim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDkimOk() (*DomainDkim, bool) {
	if o == nil || IsNil(o.Dkim) {
		return nil, false
	}
	return o.Dkim, true
}

// HasDkim returns a boolean if a field has been set.
func (o *Domain) HasDkim() bool {
	if o != nil && !IsNil(o.Dkim) {
		return true
	}

	return false
}

// SetDkim gets a reference to the given DomainDkim and assigns it to the Dkim field.
func (o *Domain) SetDkim(v DomainDkim) {
	o.Dkim = &v
}

// GetSpf returns the Spf field value if set, zero value otherwise.
func (o *Domain) GetSpf() DomainSpf {
	if o == nil || IsNil(o.Spf) {
		var ret DomainSpf
		return ret
	}
	return *o.Spf
}

// GetSpfOk returns a tuple with the Spf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSpfOk() (*DomainSpf, bool) {
	if o == nil || IsNil(o.Spf) {
		return nil, false
	}
	return o.Spf, true
}

// HasSpf returns a boolean if a field has been set.
func (o *Domain) HasSpf() bool {
	if o != nil && !IsNil(o.Spf) {
		return true
	}

	return false
}

// SetSpf gets a reference to the given DomainSpf and assigns it to the Spf field.
func (o *Domain) SetSpf(v DomainSpf) {
	o.Spf = &v
}

// GetReturnPath returns the ReturnPath field value if set, zero value otherwise.
func (o *Domain) GetReturnPath() DomainReturnPath {
	if o == nil || IsNil(o.ReturnPath) {
		var ret DomainReturnPath
		return ret
	}
	return *o.ReturnPath
}

// GetReturnPathOk returns a tuple with the ReturnPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetReturnPathOk() (*DomainReturnPath, bool) {
	if o == nil || IsNil(o.ReturnPath) {
		return nil, false
	}
	return o.ReturnPath, true
}

// HasReturnPath returns a boolean if a field has been set.
func (o *Domain) HasReturnPath() bool {
	if o != nil && !IsNil(o.ReturnPath) {
		return true
	}

	return false
}

// SetReturnPath gets a reference to the given DomainReturnPath and assigns it to the ReturnPath field.
func (o *Domain) SetReturnPath(v DomainReturnPath) {
	o.ReturnPath = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *Domain) GetTrack() DomainTrack {
	if o == nil || IsNil(o.Track) {
		var ret DomainTrack
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTrackOk() (*DomainTrack, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *Domain) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given DomainTrack and assigns it to the Track field.
func (o *Domain) SetTrack(v DomainTrack) {
	o.Track = &v
}

// GetDmarc returns the Dmarc field value if set, zero value otherwise.
func (o *Domain) GetDmarc() DomainDmarc {
	if o == nil || IsNil(o.Dmarc) {
		var ret DomainDmarc
		return ret
	}
	return *o.Dmarc
}

// GetDmarcOk returns a tuple with the Dmarc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDmarcOk() (*DomainDmarc, bool) {
	if o == nil || IsNil(o.Dmarc) {
		return nil, false
	}
	return o.Dmarc, true
}

// HasDmarc returns a boolean if a field has been set.
func (o *Domain) HasDmarc() bool {
	if o != nil && !IsNil(o.Dmarc) {
		return true
	}

	return false
}

// SetDmarc gets a reference to the given DomainDmarc and assigns it to the Dmarc field.
func (o *Domain) SetDmarc(v DomainDmarc) {
	o.Dmarc = &v
}

// GetDkimConfig returns the DkimConfig field value if set, zero value otherwise.
func (o *Domain) GetDkimConfig() string {
	if o == nil || IsNil(o.DkimConfig) {
		var ret string
		return ret
	}
	return *o.DkimConfig
}

// GetDkimConfigOk returns a tuple with the DkimConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDkimConfigOk() (*string, bool) {
	if o == nil || IsNil(o.DkimConfig) {
		return nil, false
	}
	return o.DkimConfig, true
}

// HasDkimConfig returns a boolean if a field has been set.
func (o *Domain) HasDkimConfig() bool {
	if o != nil && !IsNil(o.DkimConfig) {
		return true
	}

	return false
}

// SetDkimConfig gets a reference to the given string and assigns it to the DkimConfig field.
func (o *Domain) SetDkimConfig(v string) {
	o.DkimConfig = &v
}

// GetDkimVerified returns the DkimVerified field value if set, zero value otherwise.
func (o *Domain) GetDkimVerified() bool {
	if o == nil || IsNil(o.DkimVerified) {
		var ret bool
		return ret
	}
	return *o.DkimVerified
}

// GetDkimVerifiedOk returns a tuple with the DkimVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDkimVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.DkimVerified) {
		return nil, false
	}
	return o.DkimVerified, true
}

// HasDkimVerified returns a boolean if a field has been set.
func (o *Domain) HasDkimVerified() bool {
	if o != nil && !IsNil(o.DkimVerified) {
		return true
	}

	return false
}

// SetDkimVerified gets a reference to the given bool and assigns it to the DkimVerified field.
func (o *Domain) SetDkimVerified(v bool) {
	o.DkimVerified = &v
}

// GetSpfVerified returns the SpfVerified field value if set, zero value otherwise.
func (o *Domain) GetSpfVerified() bool {
	if o == nil || IsNil(o.SpfVerified) {
		var ret bool
		return ret
	}
	return *o.SpfVerified
}

// GetSpfVerifiedOk returns a tuple with the SpfVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSpfVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.SpfVerified) {
		return nil, false
	}
	return o.SpfVerified, true
}

// HasSpfVerified returns a boolean if a field has been set.
func (o *Domain) HasSpfVerified() bool {
	if o != nil && !IsNil(o.SpfVerified) {
		return true
	}

	return false
}

// SetSpfVerified gets a reference to the given bool and assigns it to the SpfVerified field.
func (o *Domain) SetSpfVerified(v bool) {
	o.SpfVerified = &v
}

// GetMailboxVerified returns the MailboxVerified field value if set, zero value otherwise.
func (o *Domain) GetMailboxVerified() bool {
	if o == nil || IsNil(o.MailboxVerified) {
		var ret bool
		return ret
	}
	return *o.MailboxVerified
}

// GetMailboxVerifiedOk returns a tuple with the MailboxVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetMailboxVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.MailboxVerified) {
		return nil, false
	}
	return o.MailboxVerified, true
}

// HasMailboxVerified returns a boolean if a field has been set.
func (o *Domain) HasMailboxVerified() bool {
	if o != nil && !IsNil(o.MailboxVerified) {
		return true
	}

	return false
}

// SetMailboxVerified gets a reference to the given bool and assigns it to the MailboxVerified field.
func (o *Domain) SetMailboxVerified(v bool) {
	o.MailboxVerified = &v
}

// GetDmarcVerified returns the DmarcVerified field value if set, zero value otherwise.
func (o *Domain) GetDmarcVerified() bool {
	if o == nil || IsNil(o.DmarcVerified) {
		var ret bool
		return ret
	}
	return *o.DmarcVerified
}

// GetDmarcVerifiedOk returns a tuple with the DmarcVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDmarcVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.DmarcVerified) {
		return nil, false
	}
	return o.DmarcVerified, true
}

// HasDmarcVerified returns a boolean if a field has been set.
func (o *Domain) HasDmarcVerified() bool {
	if o != nil && !IsNil(o.DmarcVerified) {
		return true
	}

	return false
}

// SetDmarcVerified gets a reference to the given bool and assigns it to the DmarcVerified field.
func (o *Domain) SetDmarcVerified(v bool) {
	o.DmarcVerified = &v
}

// GetReturnPathVerified returns the ReturnPathVerified field value if set, zero value otherwise.
func (o *Domain) GetReturnPathVerified() bool {
	if o == nil || IsNil(o.ReturnPathVerified) {
		var ret bool
		return ret
	}
	return *o.ReturnPathVerified
}

// GetReturnPathVerifiedOk returns a tuple with the ReturnPathVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetReturnPathVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnPathVerified) {
		return nil, false
	}
	return o.ReturnPathVerified, true
}

// HasReturnPathVerified returns a boolean if a field has been set.
func (o *Domain) HasReturnPathVerified() bool {
	if o != nil && !IsNil(o.ReturnPathVerified) {
		return true
	}

	return false
}

// SetReturnPathVerified gets a reference to the given bool and assigns it to the ReturnPathVerified field.
func (o *Domain) SetReturnPathVerified(v bool) {
	o.ReturnPathVerified = &v
}

// GetTrackVerified returns the TrackVerified field value if set, zero value otherwise.
func (o *Domain) GetTrackVerified() bool {
	if o == nil || IsNil(o.TrackVerified) {
		var ret bool
		return ret
	}
	return *o.TrackVerified
}

// GetTrackVerifiedOk returns a tuple with the TrackVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTrackVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackVerified) {
		return nil, false
	}
	return o.TrackVerified, true
}

// HasTrackVerified returns a boolean if a field has been set.
func (o *Domain) HasTrackVerified() bool {
	if o != nil && !IsNil(o.TrackVerified) {
		return true
	}

	return false
}

// SetTrackVerified gets a reference to the given bool and assigns it to the TrackVerified field.
func (o *Domain) SetTrackVerified(v bool) {
	o.TrackVerified = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *Domain) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *Domain) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *Domain) SetVerified(v bool) {
	o.Verified = &v
}

// GetDomainRegisteredDate returns the DomainRegisteredDate field value if set, zero value otherwise.
func (o *Domain) GetDomainRegisteredDate() string {
	if o == nil || IsNil(o.DomainRegisteredDate) {
		var ret string
		return ret
	}
	return *o.DomainRegisteredDate
}

// GetDomainRegisteredDateOk returns a tuple with the DomainRegisteredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainRegisteredDateOk() (*string, bool) {
	if o == nil || IsNil(o.DomainRegisteredDate) {
		return nil, false
	}
	return o.DomainRegisteredDate, true
}

// HasDomainRegisteredDate returns a boolean if a field has been set.
func (o *Domain) HasDomainRegisteredDate() bool {
	if o != nil && !IsNil(o.DomainRegisteredDate) {
		return true
	}

	return false
}

// SetDomainRegisteredDate gets a reference to the given string and assigns it to the DomainRegisteredDate field.
func (o *Domain) SetDomainRegisteredDate(v string) {
	o.DomainRegisteredDate = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Domain) GetCreated() int64 {
	if o == nil || IsNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Domain) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Domain) SetCreated(v int64) {
	o.Created = &v
}

// GetGptVerified returns the GptVerified field value if set, zero value otherwise.
func (o *Domain) GetGptVerified() bool {
	if o == nil || IsNil(o.GptVerified) {
		var ret bool
		return ret
	}
	return *o.GptVerified
}

// GetGptVerifiedOk returns a tuple with the GptVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetGptVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.GptVerified) {
		return nil, false
	}
	return o.GptVerified, true
}

// HasGptVerified returns a boolean if a field has been set.
func (o *Domain) HasGptVerified() bool {
	if o != nil && !IsNil(o.GptVerified) {
		return true
	}

	return false
}

// SetGptVerified gets a reference to the given bool and assigns it to the GptVerified field.
func (o *Domain) SetGptVerified(v bool) {
	o.GptVerified = &v
}

// GetGpt returns the Gpt field value if set, zero value otherwise.
func (o *Domain) GetGpt() DomainGpt {
	if o == nil || IsNil(o.Gpt) {
		var ret DomainGpt
		return ret
	}
	return *o.Gpt
}

// GetGptOk returns a tuple with the Gpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetGptOk() (*DomainGpt, bool) {
	if o == nil || IsNil(o.Gpt) {
		return nil, false
	}
	return o.Gpt, true
}

// HasGpt returns a boolean if a field has been set.
func (o *Domain) HasGpt() bool {
	if o != nil && !IsNil(o.Gpt) {
		return true
	}

	return false
}

// SetGpt gets a reference to the given DomainGpt and assigns it to the Gpt field.
func (o *Domain) SetGpt(v DomainGpt) {
	o.Gpt = &v
}

// GetDmarcFailureReason returns the DmarcFailureReason field value if set, zero value otherwise.
func (o *Domain) GetDmarcFailureReason() string {
	if o == nil || IsNil(o.DmarcFailureReason) {
		var ret string
		return ret
	}
	return *o.DmarcFailureReason
}

// GetDmarcFailureReasonOk returns a tuple with the DmarcFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDmarcFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DmarcFailureReason) {
		return nil, false
	}
	return o.DmarcFailureReason, true
}

// HasDmarcFailureReason returns a boolean if a field has been set.
func (o *Domain) HasDmarcFailureReason() bool {
	if o != nil && !IsNil(o.DmarcFailureReason) {
		return true
	}

	return false
}

// SetDmarcFailureReason gets a reference to the given string and assigns it to the DmarcFailureReason field.
func (o *Domain) SetDmarcFailureReason(v string) {
	o.DmarcFailureReason = &v
}

// GetDkimFailureReason returns the DkimFailureReason field value if set, zero value otherwise.
func (o *Domain) GetDkimFailureReason() string {
	if o == nil || IsNil(o.DkimFailureReason) {
		var ret string
		return ret
	}
	return *o.DkimFailureReason
}

// GetDkimFailureReasonOk returns a tuple with the DkimFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDkimFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DkimFailureReason) {
		return nil, false
	}
	return o.DkimFailureReason, true
}

// HasDkimFailureReason returns a boolean if a field has been set.
func (o *Domain) HasDkimFailureReason() bool {
	if o != nil && !IsNil(o.DkimFailureReason) {
		return true
	}

	return false
}

// SetDkimFailureReason gets a reference to the given string and assigns it to the DkimFailureReason field.
func (o *Domain) SetDkimFailureReason(v string) {
	o.DkimFailureReason = &v
}

// GetTrackFailureReason returns the TrackFailureReason field value if set, zero value otherwise.
func (o *Domain) GetTrackFailureReason() string {
	if o == nil || IsNil(o.TrackFailureReason) {
		var ret string
		return ret
	}
	return *o.TrackFailureReason
}

// GetTrackFailureReasonOk returns a tuple with the TrackFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTrackFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.TrackFailureReason) {
		return nil, false
	}
	return o.TrackFailureReason, true
}

// HasTrackFailureReason returns a boolean if a field has been set.
func (o *Domain) HasTrackFailureReason() bool {
	if o != nil && !IsNil(o.TrackFailureReason) {
		return true
	}

	return false
}

// SetTrackFailureReason gets a reference to the given string and assigns it to the TrackFailureReason field.
func (o *Domain) SetTrackFailureReason(v string) {
	o.TrackFailureReason = &v
}

// GetReturnPathFailureReason returns the ReturnPathFailureReason field value if set, zero value otherwise.
func (o *Domain) GetReturnPathFailureReason() string {
	if o == nil || IsNil(o.ReturnPathFailureReason) {
		var ret string
		return ret
	}
	return *o.ReturnPathFailureReason
}

// GetReturnPathFailureReasonOk returns a tuple with the ReturnPathFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetReturnPathFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnPathFailureReason) {
		return nil, false
	}
	return o.ReturnPathFailureReason, true
}

// HasReturnPathFailureReason returns a boolean if a field has been set.
func (o *Domain) HasReturnPathFailureReason() bool {
	if o != nil && !IsNil(o.ReturnPathFailureReason) {
		return true
	}

	return false
}

// SetReturnPathFailureReason gets a reference to the given string and assigns it to the ReturnPathFailureReason field.
func (o *Domain) SetReturnPathFailureReason(v string) {
	o.ReturnPathFailureReason = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Dkim) {
		toSerialize["dkim"] = o.Dkim
	}
	if !IsNil(o.Spf) {
		toSerialize["spf"] = o.Spf
	}
	if !IsNil(o.ReturnPath) {
		toSerialize["returnPath"] = o.ReturnPath
	}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	if !IsNil(o.Dmarc) {
		toSerialize["dmarc"] = o.Dmarc
	}
	if !IsNil(o.DkimConfig) {
		toSerialize["dkimConfig"] = o.DkimConfig
	}
	if !IsNil(o.DkimVerified) {
		toSerialize["dkimVerified"] = o.DkimVerified
	}
	if !IsNil(o.SpfVerified) {
		toSerialize["spfVerified"] = o.SpfVerified
	}
	if !IsNil(o.MailboxVerified) {
		toSerialize["mailboxVerified"] = o.MailboxVerified
	}
	if !IsNil(o.DmarcVerified) {
		toSerialize["dmarcVerified"] = o.DmarcVerified
	}
	if !IsNil(o.ReturnPathVerified) {
		toSerialize["returnPathVerified"] = o.ReturnPathVerified
	}
	if !IsNil(o.TrackVerified) {
		toSerialize["trackVerified"] = o.TrackVerified
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !IsNil(o.DomainRegisteredDate) {
		toSerialize["domainRegisteredDate"] = o.DomainRegisteredDate
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.GptVerified) {
		toSerialize["gptVerified"] = o.GptVerified
	}
	if !IsNil(o.Gpt) {
		toSerialize["gpt"] = o.Gpt
	}
	if !IsNil(o.DmarcFailureReason) {
		toSerialize["dmarcFailureReason"] = o.DmarcFailureReason
	}
	if !IsNil(o.DkimFailureReason) {
		toSerialize["dkimFailureReason"] = o.DkimFailureReason
	}
	if !IsNil(o.TrackFailureReason) {
		toSerialize["trackFailureReason"] = o.TrackFailureReason
	}
	if !IsNil(o.ReturnPathFailureReason) {
		toSerialize["returnPathFailureReason"] = o.ReturnPathFailureReason
	}
	return toSerialize, nil
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


