package puppetclient

// Debugf logs at info level
func (c *PuppetClient) Debugf(msg string, a ...interface{}) {
	c.clientOpts.logger.Debugf(msg, a...)
}

// Infof logs at info level
func (c *PuppetClient) Infof(msg string, a ...interface{}) {
	c.clientOpts.logger.Infof(msg, a...)
}

// Warnf logs at info level
func (c *PuppetClient) Warnf(msg string, a ...interface{}) {
	c.clientOpts.logger.Warnf(msg, a...)
}

// Errorf logs at info level
func (c *PuppetClient) Errorf(msg string, a ...interface{}) {
	c.clientOpts.logger.Errorf(msg, a...)
}
