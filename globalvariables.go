package main

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

var client *whatsmeow.Client
var container *sqlstore.Container

var number string
var message string
var messageID string
var name string
var groupID string
var startMessage string = `
Hello %s! 👋 We're glad to have you here. How can we assist you today?

Here are the commands you can use:

    /start - Get a warm welcome and list of commands (You're already here! 😊)
    /help - Get assistance with using the bot and its features.
    /verify - Check the status of your verification.
    /info - Get information about our services and policies.
    /contact - Contact support or get our contact details.
    /feedback - Send us your valuable feedback or suggestions.

Feel free to choose any command, and we'll be happy to help you! ✨
`
var helpMessage string = `
Hello! 👋 Here is a list of commands you can use to interact with me:

/start - Get a warm welcome and list of commands (You're already here! 😊)
/help - View this help message to learn more about how to use the bot
/verify - Check the status of your verification process
/info - Get information about our services, policies, and more
/contact - Contact support or get our contact details
/feedback - Share your feedback or suggestions with us

If you need any assistance, feel free to ask! I'm here to help you. ✨
`
var verifyMessage string = `*Verification Process*

Please send your IITM student email ID

Ex:- /email 23f0000000@ds.study.iitm.ac.in
`
var infoMessage string = `I am Sylph. My creator is Rudeus.`
var contactMessage string = `
📞 **Contact Support**:

If you need assistance or have any questions, feel free to get in touch with our support team:

- **Email**: support@example.com
- **Phone**: +123 456 7890 (Available 9 AM - 6 PM, Mon-Fri)
- **Live Chat**: Visit our website at [example.com](https://example.com) for live chat support.

You can also connect with us on our social media:
- **Facebook**: [facebook.com/ourpage](https://facebook.com/ourpage)
- **Twitter**: [@ourhandle](https://twitter.com/ourhandle)

We're here to help you with any queries or issues you may have! ✨
`
var feedbackMessage string = `
📝 **We Value Your Feedback**:

Your feedback helps us improve our service. We'd love to hear your thoughts!

Please let us know what you think about our service, how we can improve, or if you encountered any issues. You can provide feedback by:

1. **Replying to this message** with your thoughts.
2. **Filling out our feedback form**: [feedback-form-link](https://example.com/feedback)

If you're facing any problems or have suggestions, we're all ears! Thank you for helping us serve you better. 😊

Your opinion matters to us! ✨
`

var otp int
