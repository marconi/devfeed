(function() {
  define(["devfeed"], function(Devfeed) {
    Devfeed.module("Common.Utils", function(Utils, Devfeed, Backbone, Marionette, $, _) {
      return Utils.SmallSpin = {
        lines: 8,
        length: 4,
        width: 3,
        radius: 5,
        corners: 1
      };
    });
    return Devfeed.Common.Utils;
  });

}).call(this);
