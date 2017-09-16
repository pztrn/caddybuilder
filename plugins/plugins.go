// Caddybuilder - a friendly tool to build Caddy executable.
// Copyright (c) 2017, Stanislav N. aka pztrn <pztrn at pztrn dot name>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// File "plugins.go" contains controlling struct for plugins, which is
// responsible for all actions with plugins (initialization, plugins
// context initialization and so on.)
package plugins

import (
	// local
	"github.com/pztrn/caddybuilder/plugins/context"
	"github.com/pztrn/caddybuilder/plugins/handler"

	// plugins
	"github.com/pztrn/caddybuilder/plugins/authz"
	"github.com/pztrn/caddybuilder/plugins/awses"
	"github.com/pztrn/caddybuilder/plugins/awslambda"
	"github.com/pztrn/caddybuilder/plugins/cache"
	"github.com/pztrn/caddybuilder/plugins/cgi"
	"github.com/pztrn/caddybuilder/plugins/cors"
	"github.com/pztrn/caddybuilder/plugins/datadog"
	"github.com/pztrn/caddybuilder/plugins/expires"
	"github.com/pztrn/caddybuilder/plugins/filemanager"
	"github.com/pztrn/caddybuilder/plugins/filter"
	"github.com/pztrn/caddybuilder/plugins/forwardproxy"
	"github.com/pztrn/caddybuilder/plugins/git"
	"github.com/pztrn/caddybuilder/plugins/gopkg"
	"github.com/pztrn/caddybuilder/plugins/grpc"
	"github.com/pztrn/caddybuilder/plugins/ipfilter"
	"github.com/pztrn/caddybuilder/plugins/jsonp"
	"github.com/pztrn/caddybuilder/plugins/jwt"
	"github.com/pztrn/caddybuilder/plugins/locale"
	"github.com/pztrn/caddybuilder/plugins/mailout"
	"github.com/pztrn/caddybuilder/plugins/minify"
	"github.com/pztrn/caddybuilder/plugins/nobots"
	"github.com/pztrn/caddybuilder/plugins/prometheus"
	"github.com/pztrn/caddybuilder/plugins/proxyprotocol"
	"github.com/pztrn/caddybuilder/plugins/ratelimit"
	"github.com/pztrn/caddybuilder/plugins/realip"
	"github.com/pztrn/caddybuilder/plugins/reauth"
	"github.com/pztrn/caddybuilder/plugins/restic"
	"github.com/pztrn/caddybuilder/plugins/upload"
	"github.com/pztrn/caddybuilder/plugins/webdav"
)

type Plugins struct {
	// Plugins context.
	pluginContext *plugincontext.PluginContext
	// List with initialized plugins.
	// Name is a plugin name, value - plugin interface.
	PluginsList map[string]pluginhandler.PluginHandler
}

// Initializes plugins controller.
func (p *Plugins) Initialize() {
	p.PluginsList = make(map[string]pluginhandler.PluginHandler)
	p.initializePluginContext()
	p.initializePlugins()
}

// Initializes plugins context, which will be passed to plugins.
func (p *Plugins) initializePluginContext() {
	log.Print("Initializing plugins context...")
	p.pluginContext = &plugincontext.PluginContext{}
	p.pluginContext.CmdWorker = cw
	p.pluginContext.Flags = flags
	p.pluginContext.Log = log
}

// Initializes plugins.
func (p *Plugins) initializePlugins() {
	log.Print("Initializing plugins...")

	authz_raw := authz.New(p.pluginContext)
	authz_interface := pluginhandler.PluginHandler(authz_raw)
	authz_interface.Initialize()
	p.PluginsList[authz_interface.GetPluginName()] = authz_interface

	awses_raw := awses.New(p.pluginContext)
	awses_interface := pluginhandler.PluginHandler(awses_raw)
	awses_interface.Initialize()
	p.PluginsList[awses_interface.GetPluginName()] = awses_interface

	awslambda_raw := awslambda.New(p.pluginContext)
	awslambda_interface := pluginhandler.PluginHandler(awslambda_raw)
	awslambda_interface.Initialize()
	p.PluginsList[awslambda_interface.GetPluginName()] = awslambda_interface

	cache_raw := cache.New(p.pluginContext)
	cache_interface := pluginhandler.PluginHandler(cache_raw)
	cache_interface.Initialize()
	p.PluginsList[cache_interface.GetPluginName()] = cache_interface

	cgi_raw := cgi.New(p.pluginContext)
	cgi_interface := pluginhandler.PluginHandler(cgi_raw)
	cgi_interface.Initialize()
	p.PluginsList[cgi_interface.GetPluginName()] = cgi_interface

	cors_raw := cors.New(p.pluginContext)
	cors_interface := pluginhandler.PluginHandler(cors_raw)
	cors_interface.Initialize()
	p.PluginsList[cors_interface.GetPluginName()] = cors_interface

	datadog_raw := datadog.New(p.pluginContext)
	datadog_interface := pluginhandler.PluginHandler(datadog_raw)
	datadog_interface.Initialize()
	p.PluginsList[datadog_interface.GetPluginName()] = datadog_interface

	expires_raw := expires.New(p.pluginContext)
	expires_interface := pluginhandler.PluginHandler(expires_raw)
	expires_interface.Initialize()
	p.PluginsList[expires_interface.GetPluginName()] = expires_interface

	filemanager_raw := filemanager.New(p.pluginContext)
	filemanager_interface := pluginhandler.PluginHandler(filemanager_raw)
	filemanager_interface.Initialize()
	p.PluginsList[filemanager_interface.GetPluginName()] = filemanager_interface

	filter_raw := filter.New(p.pluginContext)
	filter_interface := pluginhandler.PluginHandler(filter_raw)
	filter_interface.Initialize()
	p.PluginsList[filter_interface.GetPluginName()] = filter_interface

	forwardproxy_raw := forwardproxy.New(p.pluginContext)
	forwardproxy_interface := pluginhandler.PluginHandler(forwardproxy_raw)
	forwardproxy_interface.Initialize()
	p.PluginsList[forwardproxy_interface.GetPluginName()] = forwardproxy_interface

	git_raw := git.New(p.pluginContext)
	git_interface := pluginhandler.PluginHandler(git_raw)
	git_interface.Initialize()
	p.PluginsList[git_interface.GetPluginName()] = git_interface

	gopkg_raw := gopkg.New(p.pluginContext)
	gopkg_interface := pluginhandler.PluginHandler(gopkg_raw)
	gopkg_interface.Initialize()
	p.PluginsList[gopkg_interface.GetPluginName()] = gopkg_interface

	grpc_raw := grpc.New(p.pluginContext)
	grpc_interface := pluginhandler.PluginHandler(grpc_raw)
	grpc_interface.Initialize()
	p.PluginsList[grpc_interface.GetPluginName()] = grpc_interface

	ipfilter_raw := ipfilter.New(p.pluginContext)
	ipfilter_interface := pluginhandler.PluginHandler(ipfilter_raw)
	ipfilter_interface.Initialize()
	p.PluginsList[ipfilter_interface.GetPluginName()] = ipfilter_interface

	jsonp_raw := jsonp.New(p.pluginContext)
	jsonp_interface := pluginhandler.PluginHandler(jsonp_raw)
	jsonp_interface.Initialize()
	p.PluginsList[jsonp_interface.GetPluginName()] = jsonp_interface

	jwt_raw := jwt.New(p.pluginContext)
	jwt_interface := pluginhandler.PluginHandler(jwt_raw)
	jwt_interface.Initialize()
	p.PluginsList[jwt_interface.GetPluginName()] = jwt_interface

	locale_raw := locale.New(p.pluginContext)
	locale_interface := pluginhandler.PluginHandler(locale_raw)
	locale_interface.Initialize()
	p.PluginsList[locale_interface.GetPluginName()] = locale_interface

	mailout_raw := mailout.New(p.pluginContext)
	mailout_interface := pluginhandler.PluginHandler(mailout_raw)
	mailout_interface.Initialize()
	p.PluginsList[mailout_interface.GetPluginName()] = mailout_interface

	minify_raw := minify.New(p.pluginContext)
	minify_interface := pluginhandler.PluginHandler(minify_raw)
	minify_interface.Initialize()
	p.PluginsList[minify_interface.GetPluginName()] = minify_interface

	nobots_raw := nobots.New(p.pluginContext)
	nobots_interface := pluginhandler.PluginHandler(nobots_raw)
	nobots_interface.Initialize()
	p.PluginsList[nobots_interface.GetPluginName()] = nobots_interface

	prometheus_raw := prometheus.New(p.pluginContext)
	prometheus_interface := pluginhandler.PluginHandler(prometheus_raw)
	prometheus_interface.Initialize()
	p.PluginsList[prometheus_interface.GetPluginName()] = prometheus_interface

	proxyprotocol_raw := proxyprotocol.New(p.pluginContext)
	proxyprotocol_interface := pluginhandler.PluginHandler(proxyprotocol_raw)
	proxyprotocol_interface.Initialize()
	p.PluginsList[proxyprotocol_interface.GetPluginName()] = proxyprotocol_interface

	ratelimit_raw := ratelimit.New(p.pluginContext)
	ratelimit_interface := pluginhandler.PluginHandler(ratelimit_raw)
	ratelimit_interface.Initialize()
	p.PluginsList[ratelimit_interface.GetPluginName()] = ratelimit_interface

	realip_raw := realip.New(p.pluginContext)
	realip_interface := pluginhandler.PluginHandler(realip_raw)
	realip_interface.Initialize()
	p.PluginsList[realip_interface.GetPluginName()] = realip_interface

	reauth_raw := reauth.New(p.pluginContext)
	reauth_interface := pluginhandler.PluginHandler(reauth_raw)
	reauth_interface.Initialize()
	p.PluginsList[reauth_interface.GetPluginName()] = reauth_interface

	restic_raw := restic.New(p.pluginContext)
	restic_interface := pluginhandler.PluginHandler(restic_raw)
	restic_interface.Initialize()
	p.PluginsList[restic_interface.GetPluginName()] = restic_interface

	upload_raw := upload.New(p.pluginContext)
	upload_interface := pluginhandler.PluginHandler(upload_raw)
	upload_interface.Initialize()
	p.PluginsList[upload_interface.GetPluginName()] = upload_interface

	webdav_raw := webdav.New(p.pluginContext)
	webdav_interface := pluginhandler.PluginHandler(webdav_raw)
	webdav_interface.Initialize()
	p.PluginsList[webdav_interface.GetPluginName()] = webdav_interface
}
