!function(e){"function"==typeof define&&define.amd?define(["jquery"],e):e("object"==typeof exports?require("jquery"):jQuery)}(function(e){"use strict";function t(i,n){this.$element=e(i),this.options=e.extend({},t.DEFAULTS,e.isPlainObject(n)&&n),this.init()}var i="qor.collection.sortable",n="enable."+i,o="disable."+i,s="click."+i,r=".qor-sortable__item",d=".qor-sortable__button-change",a=".qor-sortable__button-done",l=".qor-sortable__button-add",c=".qor-sortable__button-delete",f=".qor-sortable__button-move",h=".qor-sortable__action",u=".qor-sortable__action-position",m=".is-delete",b="sortable-collection-loaded";return t.prototype={constructor:t,init:function(){this.bind(),this.initItemOrder()},bind:function(){this.$element.on(s,e.proxy(this.click,this))},unbind:function(){this.$element.off(s,this.click)},initItemOrder:function(i){var n=this.$element.find(r).filter(":visible").not(m);if(n.size()){var o,s,d=n.find(h).find(u),a={},l=n.size(),c=n.first().html();d.size()&&d.remove(),s||(o=c.match(/(\w+)\="(\S*\[\d+\]\S*)"/),o.length&&(o=o[2],s=o.match(/^(\S*)\[(\d+)\]([^\[\]]*)$/),s.length&&(s=s[1]))),n.each(function(n){var o=e(this),r=o.find(h);a.isSelected=!1,a.itemTotal=l,a.itemIndex=n+1,r.prepend('<select class="qor-sortable__action-position"></select>');for(var d=1;l>=d;d++)a.index=d,a.itemIndex==d?a.isSelected=!0:a.isSelected=!1,r.find("select").append(window.Mustache.render(t.OPTION_HTML,a));if(i){var c,f,u=e(this).find('[name^="'+s+'"]');u.size()&&u.each(function(){c=e(this).prop("name"),f="["+a.itemIndex+"]",c=c.replace(/\[\d+\]/,f),e(this).prop("name",c)})}o.data(a)})}},moveItem:function(t){var i,n,o=t.closest(r),s=o.data().itemIndex,d=o.find(u).val();d!=s&&(i=1==d?1:s>d?d-1:d,n=e(r).filter(function(){return e(this).data().itemIndex==i}),1==d?n.before(o.fadeOut("slow").fadeIn("slow")):n.after(o.fadeOut("slow").fadeIn("slow")),this.initItemOrder(!0))},click:function(t){var i=e(t.target),n=this.$element,o=n.find(r).filter(":visible").not(m);i.is(f)&&this.moveItem(i),i.is(a)&&(i.hide(),n.find(h).hide(),n.find(d).show(),n.find(l).show(),n.find(c).show()),i.is(d)&&o.size()&&(i.hide(),n.find(a).show(),n.find(h).show(),n.find(l).hide(),n.find(c).hide(),this.initItemOrder())},destroy:function(){this.unbind(),this.$element.removeData(i)}},t.DEFAULTS={},t.OPTION_HTML='<option value="[[index]]" [[#isSelected]]selected[[/isSelected]]>[[index]] of [[itemTotal]]</option>',t.plugin=function(n){return this.each(function(){var o,s=e(this),r=s.data(i);if(!r){if(/destroy/.test(n))return;s.data(i,r=new t(this,n))}"string"==typeof n&&e.isFunction(o=r[n])&&o.apply(r)})},e(function(){var i='[data-toggle="qor.collection.sortable"]';e("body").data(b)||(e("body").data(b,!0),e(document).on(o,function(n){t.plugin.call(e(i,n.target),"destroy")}).on(n,function(n){t.plugin.call(e(i,n.target))}).triggerHandler(n))}),t});