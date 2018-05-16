package controllers

import (
	//	"encoding/json"
	//	"fmt"
	//	"strconv"
	//	"bufio"
	//	"encoding/json"
	//	"encoding/base64"
	//	"io/ioutil"

	//	"strings"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
	//	"github.com/gorilla/schema"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "numberview.tpl"
}
