(function() {
  _.mixin({
    compactObject: function(o) {
      _.each(o, function(v, k) {
        if (!v) {
          return delete o[k];
        }
      });
      return o;
    }
  });

}).call(this);
