package adacorebase

// VERSION - adacore version
const VERSION = "0.2.38"

// BigMsgLength -if msg length >= BigMsgLength, the message is big message
const BigMsgLength = 4*1024*1024 - 1024

// DefaultMaxExpireTime - max expire time in seconds
const DefaultMaxExpireTime = 24 * 60 * 60

// DefaultIsAllowTemplateData - Whether to allow templatedata
const DefaultIsAllowTemplateData = false

// DefaultTemplate - This is default template available for this role.
const DefaultTemplate = "default"

// DefaultResNums - This is the amount of resources available for this role
const DefaultResNums = 0

// DefaultBaseURL - default baseurl
const DefaultBaseURL = "http://127.0.0.1/"

// DefaultTemplatesPath - default templates file path
const DefaultTemplatesPath = "./templates"
