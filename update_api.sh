#!/bin/bash

goctl api go -api gateway/gateway.api -dir gateway
goctl api ts -api gateway/gateway.api -dir ~/H5Projects/video-site/src/api -webapi "../utils/http"