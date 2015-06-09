package commands

import "fmt"

func cmdStatus(args commandArgs, c *Conn) {
	mailbox, err := c.user.MailboxByName(args.Arg(0))
	if err != nil {
		c.writeResponse(args.ID(), "NO "+err.Error())
		return
	}

	c.writeResponse("", fmt.Sprintf("STATUS %s (UIDNEXT %d UNSEEN %d)",
		mailbox.Name(), mailbox.NextUID(), mailbox.Unseen()))
	c.writeResponse(args.ID(), "OK STATUS Completed")
}