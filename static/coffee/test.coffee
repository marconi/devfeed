require.config
  baseUrl: 'js'

  # NOTE: remove on production, only for development
  # to prevent required script from being cached.
  urlArgs: "bust=" + (new Date()).getTime()

  paths:
    # vendors
    json2: "vendors/json2"
    jquery: "vendors/jquery"
    jquery_spin: "vendors/jquery.spin"
    underscore: "vendors/underscore"
    backbone: "vendors/backbone"
    syphon: "vendors/backbone.syphon"
    spin: "vendors/spin"
    marionette: "vendors/backbone.marionette"
    tpl: "vendors/tpl"

    # foundation
    foundation: "vendors/foundation/foundation"
    foundation_dropdown: "vendors/foundation/foundation.dropdown"
    foundation_alerts: "vendors/foundation/foundation.alerts"

    # apps
    devfeed: "devfeed"
    user_session: "entities/user_session"
    alert: "entities/alert"

    # common app
    common_view: "apps/common/common_view"

    # auth app
    auth_app: "apps/auth/auth_app"
    login_controller: "apps/auth/login/login_controller"
    login_view: "apps/auth/login/login_view"
    logout_controller: "apps/auth/logout/logout_controller"

    # header app
    header_app: "apps/header/header_app"
    header_controller: "apps/header/common/header_controller"
    header_view: "apps/header/common/header_view"

    # project app
    project_app: "apps/project/project_app"
    project_list_controller: "apps/project/list/list_controller"

    # specs
    user_session_spec: "tests/specs/user_session.spec"
    auth_app_spec: "tests/specs/auth_app.spec"
    header_app_spec: "tests/specs/header_app.spec"

  shim:
    jquery:
      exports: "$"
    jquery_spin:
      deps: ["jquery", "spin"]
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

    # foundation
    foundation:
      deps: ["jquery"]
    foundation_dropdown:
      deps: ["foundation"]
    foundation_alerts:
      deps: ["foundation"]

    # apps
    devfeed:
      deps: ["foundation_dropdown", "foundation_alerts"]

define ["user_session_spec", "auth_app_spec", "header_app_spec"], ->
  mochaPhantomJS.run()
