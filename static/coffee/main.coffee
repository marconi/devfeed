require.config
  baseUrl: "/static/js"

  # NOTE: remove on production, only for development
  # to prevent required script from being cached.
  urlArgs: "bust=" + (new Date()).getTime()

  pragmasOnSave:
    excludeTpl: true

  paths:
    # vendors
    json2: "vendors/json2"
    jquery: "vendors/jquery"
    jquery_spin: "vendors/jquery.spin"
    jquery_livefilter: "vendors/jquery.fastlivefilter"
    underscore: "vendors/underscore"
    backbone: "vendors/backbone"
    syphon: "vendors/backbone.syphon"
    spin: "vendors/spin"
    marionette: "vendors/backbone.marionette"
    validation: "vendors/backbone.validation"
    tpl: "vendors/tpl"
    notification: "vendors/mininotification"
    golem: "vendors/golem"

    # foundation
    foundation: "vendors/foundation/foundation"
    foundation_dropdown: "vendors/foundation/foundation.dropdown"
    foundation_alerts: "vendors/foundation/foundation.alerts"
    foundation_section: "vendors/foundation/foundation.section"

    # utils
    utils: "utils"

    # apps
    devfeed: "devfeed"

    # entities
    user_session: "entities/user_session"
    alert: "entities/alert"
    project: "entities/project"
    websocket: "entities/websocket"
    chat: "entities/chat"

    # common app
    common_view: "apps/common/common_view"
    common_utils: "apps/common/common_utils"

    # auth app
    auth_app: "apps/auth/auth_app"
    login_controller: "apps/auth/login/login_controller"
    login_view: "apps/auth/login/login_view"
    register_controller: "apps/auth/register/register_controller"
    register_view: "apps/auth/register/register_view"
    activation_controller: "apps/auth/activation/activation_controller"
    activation_view: "apps/auth/activation/activation_view"
    password_controller: "apps/auth/password/password_controller"
    password_view: "apps/auth/password/password_view"

    # header app
    header_app: "apps/header/header_app"
    header_controller: "apps/header/common/header_controller"
    header_view: "apps/header/common/header_view"

    # project app
    project_app: "apps/project/project_app"
    project_list_controller: "apps/project/list/list_controller"
    project_list_view: "apps/project/list/list_view"
    project_show_controller: "apps/project/show/show_controller"
    project_show_view: "apps/project/show/show_view"

    # settings app
    settings_app: "apps/settings/settings_app"
    settings_config_controller: "apps/settings/config/config_controller"
    settings_config_view: "apps/settings/config/config_view"

  shim:
    jquery:
      exports: "$"
    jquery_spin:
      deps: ["jquery", "spin"]
    jquery_livefilter:
      deps: ["jquery"]
    underscore:
      exports: "_"
    backbone:
      deps: ["jquery", "json2", "underscore"]
      exports: "Backbone"
    syphon:
      deps: ["backbone"]
    marionette:
      deps: ["backbone"]
      exports: "Marionette"
    validation:
      deps: ["backbone"]
    notification:
      deps: ["jquery"]
    websocket:
      deps: ["jquery", "golem"]

    # foundation
    foundation:
      deps: ["jquery"]
    foundation_dropdown:
      deps: ["foundation"]
    foundation_alerts:
      deps: ["foundation"]
    foundation_section:
      deps: ["foundation"]

    # utils
    utils:
      deps: ["underscore"]

    # apps
    devfeed:
      deps: ["validation", "foundation_dropdown", "foundation_alerts", "foundation_section", "notification"]

define [
  "devfeed",
  "websocket",
  "utils",
  "user_session",
  "header_app",
  "auth_app",
  "project_app",
  "settings_app"
], (Devfeed) ->

  Devfeed.start()
