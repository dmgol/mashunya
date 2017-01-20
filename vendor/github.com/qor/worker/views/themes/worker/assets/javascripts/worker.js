!function(r){"function"==typeof define&&define.amd?define(["jquery"],r):r("object"==typeof exports?require("jquery"):jQuery)}(function(r){"use strict";function e(o,t){this.$element=r(o),this.options=r.extend({},e.DEFAULTS,r.isPlainObject(t)&&t),this.init()}var o="qor.worker",t="enable."+o,s="disable."+o,n="click."+o,i=".qor-worker--new",d=".qor-worker--show-errors",a=".qor-worker-form",l=".qor-worker-form-list",u=".qor-worker--progress",c=".qor-worker-form--show",f=".qor-button--back",h=".qor-js-table",p=".is-selected";return e.prototype={constructor:e,init:function(){this.bind(),this.formOpened=!1,r(u).size()&&r.fn.qorSliderAfterShow.updateWorkerProgress(),r(c).size()&&(this.formOpened=!0)},bind:function(){this.$element.on(n,d,r.proxy(this.showError,this)).on(n,i,r.proxy(this.showForm,this)).on(n,f,r.proxy(this.hideForm,this))},unbind:function(){this.$element.off(n,d,this.showError,this).off(n,i,this.showForm,this).off(n,f,this.hideForm,this)},showError:function(o){o.preventDefault();var t=r(e.POPOVERTEMPLATE).appendTo("body"),s=r("tr.is-selected .qor-button--edit").attr("href");t.qorModal("show"),r.ajax({url:s,method:"GET",dataType:"html",processData:!1,contentType:!1}).done(function(e){var o=r(e).find(".qor-form-container"),s=o.find(".workers-error-output");s&&s.appendTo(t.find("#qor-worker-errors"))})},hideForm:function(e){e.preventDefault();var o=this.$element,t=o.find(a).find(">li");t.show().removeClass("current").find("form").addClass("hidden"),r(f).addClass("hidden"),r(l).show(),this.formOpened=!1,window.onbeforeunload=null,r.fn.qorSlideoutBeforeHide=null},showForm:function(e){var o=r(e.target);if(e.preventDefault(),!this.formOpened){var t=o.closest("li"),s=o.closest(a),n=o.closest(l),i=s.find(">li");i.hide().removeClass("current"),t.addClass("current").show(),r(f).removeClass("hidden"),t.find(l).hide(),n.show().find("form").removeClass("hidden"),this.formOpened=!0}},destroy:function(){this.unbind(),e.getWorkerProgressIntervId&&window.clearInterval(e.getWorkerProgressIntervId),r.fn.qorSliderAfterShow.updateWorkerProgress=null}},e.DEFAULTS={},e.POPOVERTEMPLATE='<div class="qor-modal fade qor-modal--worker-errors" tabindex="-1" role="dialog" aria-hidden="true"><div class="mdl-card mdl-shadow--2dp" role="document"><div class="mdl-card__title"><h2 class="mdl-card__title-text">Process Errors</h2></div><div class="mdl-card__supporting-text" id="qor-worker-errors"></div><div class="mdl-card__actions"><a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" data-dismiss="modal">close</a></div></div></div>',e.plugin=function(t){return this.each(function(){var s,n=r(this),i=n.data(o);if(!i){if(/destroy/.test(t))return;n.data(o,i=new e(this,t))}"string"==typeof t&&r.isFunction(s=i[t])&&s.apply(i)})},r.fn.qorSliderAfterShow.updateWorkerProgress=function(r){e.getWorkerProgressIntervId=window.setInterval(e.updateWorkerProgress,1e3,r)},e.updateTableStatus=function(e){var o=r(h).find(p),t=r(u).data().statusName;o.find('td[data-heading="'+t+'"]').find(".qor-table__content").html(e)},e.isScrollToBottom=function(r){return r.clientHeight+r.scrollTop===r.scrollHeight},e.updateWorkerProgress=function(o){var t=o,s=r(".workers-log-output"),n=r(".qor-worker--progress-value"),i=r(".qor-worker--progress-status"),d=r(u),a=r(h).find(p),l=["killed","exception","cancelled","scheduled"];if(d.size())var c=d.data();if(a.size()&&c&&c.statusName)var f=a.find('td[data-heading="'+c.statusName+'"]').find(".qor-table__content").html();return d.size()&&d.size()&&-1==l.indexOf(c.status)?c.progress>=100?(window.clearInterval(e.getWorkerProgressIntervId),document.querySelector("#qor-worker--progress").MaterialProgress.setProgress(100),e.updateTableStatus(c.status),r(".qor-workers-abort").addClass("hidden"),void r(".qor-workers-rerun").removeClass("hidden")):void r.ajax({url:t,method:"GET",dataType:"html",processData:!1,contentType:!1}).done(function(o){var t=r(o),d=t.find(u).data(),a=d.progress,l=d.status;n.html(a),i.html(l),document.querySelector("#qor-worker--progress").MaterialProgress.setProgress(a);var c,h=r.trim(s.html()),p=r.trim(t.find(".workers-log-output").html());p!=h&&(c=p.replace(h,""),e.isScrollToBottom(s[0])?s.append(c).scrollTop(s[0].scrollHeight):s.append(c)),f!=l&&e.updateTableStatus(l),a>=100&&(window.clearInterval(e.getWorkerProgressIntervId),r(".qor-workers-abort").addClass("hidden"),r(".qor-workers-rerun").removeClass("hidden"))}):void window.clearInterval(e.getWorkerProgressIntervId)},r(function(){var o='[data-toggle="qor.workers"]';r(document).on(s,function(t){e.plugin.call(r(o,t.target),"destroy")}).on(t,function(t){e.plugin.call(r(o,t.target))}).triggerHandler(t)}),e});