webpackJsonp([96],{266:function(n,e,t){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var i=t(743),a=t(744),r=t(0),l=r(i.a,a.a,!1,null,null,null);e.default=l.exports},441:function(n,e,t){"use strict";function i(n){t(442)}var a=t(444),r=t(445),l=t(0),o=i,s=l(a.a,r.a,!1,o,null,null);e.a=s.exports},442:function(n,e,t){var i=t(443);"string"==typeof i&&(i=[[n.i,i,""]]),i.locals&&(n.exports=i.locals);t(262)("1ef67b9a",i,!0)},443:function(n,e,t){e=n.exports=t(261)(),e.push([n.i,'/**\n* actionsheet\n*/\n/**\n* datetime\n*/\n/**\n* tabbar\n*/\n/**\n* tab\n*/\n/**\n* dialog\n*/\n/**\n* x-number\n*/\n/**\n* checkbox\n*/\n/**\n* check-icon\n*/\n/**\n* Cell\n*/\n/**\n* Mask\n*/\n/**\n* Range\n*/\n/**\n* Tabbar\n*/\n/**\n* Header\n*/\n/**\n* Timeline\n*/\n/**\n* Switch\n*/\n/**\n* Button\n*/\n/**\n* swipeout\n*/\n/**\n* Cell\n*/\n/**\n* Badge\n*/\n/**\n* Popover\n*/\n/**\n* Button tab\n*/\n/* alias */\n/**\n* Swiper\n*/\n/**\n* checklist\n*/\n/**\n* popup-picker\n*/\n/**\n* popup\n*/\n/**\n* popup-header\n*/\n/**\n* form-preview\n*/\n/**\n* load-more\n*/\n/**\n* sticky\n*/\n/**\n* group\n*/\n/**\n* toast\n*/\n/**\n* icon\n*/\n/**\n* calendar\n*/\n/**\n* week-calendar\n*/\n/**\n* search\n*/\n/**\n* radio\n*/\n/**\n* loadmore\n*/\n.weui-label {\n  display: block;\n  width: 105px;\n  word-wrap: break-word;\n  word-break: break-all;\n}\n.weui-input {\n  width: 100%;\n  border: 0;\n  outline: 0;\n  -webkit-appearance: none;\n  background-color: transparent;\n  font-size: inherit;\n  color: inherit;\n  height: 1.41176471em;\n  line-height: 1.41176471;\n}\n.weui-input::-webkit-outer-spin-button,\n.weui-input::-webkit-inner-spin-button {\n  -webkit-appearance: none;\n  margin: 0;\n}\n.weui-textarea {\n  display: block;\n  border: 0;\n  resize: none;\n  width: 100%;\n  color: inherit;\n  font-size: 1em;\n  line-height: inherit;\n  outline: 0;\n}\n.weui-textarea-counter {\n  color: #B2B2B2;\n  text-align: right;\n}\n.weui-cell_warn .weui-textarea-counter {\n  color: #E64340;\n}\n.weui-toptips {\n  display: none;\n  position: fixed;\n  -webkit-transform: translateZ(0);\n          transform: translateZ(0);\n  top: 0;\n  left: 0;\n  right: 0;\n  padding: 5px;\n  font-size: 14px;\n  text-align: center;\n  color: #FFF;\n  z-index: 5000;\n  word-wrap: break-word;\n  word-break: break-all;\n}\n.weui-toptips_warn {\n  background-color: #E64340;\n}\n.weui-cells_form .weui-cell__ft {\n  font-size: 0;\n}\n.weui-cells_form .weui-icon-warn {\n  display: none;\n}\n.weui-cells_form input,\n.weui-cells_form textarea,\n.weui-cells_form label[for] {\n  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);\n}\n.weui-cell_warn {\n  color: #E64340;\n}\n.weui-cell_warn .weui-icon-warn {\n  display: inline-block;\n}\n.weui-cell_switch {\n  padding-top: 6px;\n  padding-bottom: 6px;\n}\n.weui-switch {\n  -webkit-appearance: none;\n          appearance: none;\n}\n.weui-switch,\n.weui-switch-cp__box {\n  position: relative;\n  width: 52px;\n  height: 32px;\n  border: 1px solid #DFDFDF;\n  outline: 0;\n  border-radius: 16px;\n  box-sizing: border-box;\n  background-color: #DFDFDF;\n  -webkit-transition: background-color 0.1s, border 0.1s;\n  transition: background-color 0.1s, border 0.1s;\n}\n.weui-switch:before,\n.weui-switch-cp__box:before {\n  content: " ";\n  position: absolute;\n  top: 0;\n  left: 0;\n  width: 50px;\n  height: 30px;\n  border-radius: 15px;\n  background-color: #FDFDFD;\n  -webkit-transition: -webkit-transform 0.35s cubic-bezier(0.45, 1, 0.4, 1);\n  transition: -webkit-transform 0.35s cubic-bezier(0.45, 1, 0.4, 1);\n  transition: transform 0.35s cubic-bezier(0.45, 1, 0.4, 1);\n  transition: transform 0.35s cubic-bezier(0.45, 1, 0.4, 1), -webkit-transform 0.35s cubic-bezier(0.45, 1, 0.4, 1);\n}\n.weui-switch:after,\n.weui-switch-cp__box:after {\n  content: " ";\n  position: absolute;\n  top: 0;\n  left: 0;\n  width: 30px;\n  height: 30px;\n  border-radius: 15px;\n  background-color: #FFFFFF;\n  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);\n  -webkit-transition: -webkit-transform 0.35s cubic-bezier(0.4, 0.4, 0.25, 1.35);\n  transition: -webkit-transform 0.35s cubic-bezier(0.4, 0.4, 0.25, 1.35);\n  transition: transform 0.35s cubic-bezier(0.4, 0.4, 0.25, 1.35);\n  transition: transform 0.35s cubic-bezier(0.4, 0.4, 0.25, 1.35), -webkit-transform 0.35s cubic-bezier(0.4, 0.4, 0.25, 1.35);\n}\n.weui-switch:checked,\n.weui-switch-cp__input:checked ~ .weui-switch-cp__box {\n  /** vux **/\n  border-color: #ffe26d;\n  background-color: #ffe26d;\n  /** end vux **/\n}\n.weui-switch:checked:before,\n.weui-switch-cp__input:checked ~ .weui-switch-cp__box:before {\n  -webkit-transform: scale(0);\n          transform: scale(0);\n}\n.weui-switch:checked:after,\n.weui-switch-cp__input:checked ~ .weui-switch-cp__box:after {\n  -webkit-transform: translateX(20px);\n          transform: translateX(20px);\n}\n.weui-switch-cp__input {\n  position: absolute;\n  left: -9999px;\n}\n.weui-switch-cp__box {\n  display: block;\n}\n.weui-cell_switch .weui-cell__ft {\n  font-size: 0;\n  position: relative;\n}\ninput.weui-switch[disabled] {\n  opacity: 0.6;\n}\n.vux-x-switch.weui-cell_switch {\n  padding-top: 6px;\n  padding-bottom: 6px;\n}\n.vux-x-switch-overlay {\n  width: 60px;\n  height: 50px;\n  position: absolute;\n  right: 0;\n  top: 0;\n  opacity: 0;\n}\n',""])},444:function(n,e,t){"use strict";var i=t(41),a=t(42),r=t.n(a);e.a={name:"x-switch",components:{InlineDesc:i.a},computed:{labelStyle:function(){var n=/<\/?[^>]*>/.test(this.title),e=Math.min(n?5:this.title.length+1,14)+"em";return r()({display:"block",width:this.$parent.labelWidth||e,textAlign:this.$parent.labelAlign})},labelClass:function(){return{"vux-cell-justify":this.$parent&&"justify"===this.$parent.labelAlign}}},methods:{onClick:function(){this.$emit("on-click",!this.currentValue,this.currentValue)},toBoolean:function(n){if(this.valueMap){return 1===this.valueMap.indexOf(n)}return n},toRaw:function(n){return this.valueMap?this.valueMap[n?1:0]:n}},props:{title:{type:String,required:!0},disabled:Boolean,value:{type:[Boolean,String,Number],default:!1},inlineDesc:[String,Boolean,Number],preventDefault:Boolean,valueMap:{type:Array,default:function(){return[!1,!0]}}},data:function(){return{currentValue:this.toBoolean(this.value)}},watch:{currentValue:function(n){var e=this.toRaw(n);this.$emit("input",e),this.$emit("on-change",e)},value:function(n){this.currentValue=this.toBoolean(n)}}}},445:function(n,e,t){"use strict";var i=function(){var n=this,e=n.$createElement,t=n._self._c||e;return t("div",{staticClass:"vux-x-switch weui-cell weui-cell_switch"},[t("div",{staticClass:"weui-cell__bd"},[t("label",{staticClass:"weui-label",class:n.labelClass,style:n.labelStyle,domProps:{innerHTML:n._s(n.title)}}),n._v(" "),n.inlineDesc?t("inline-desc",[n._v(n._s(n.inlineDesc))]):n._e()],1),n._v(" "),t("div",{staticClass:"weui-cell__ft"},[t("input",{directives:[{name:"model",rawName:"v-model",value:n.currentValue,expression:"currentValue"}],staticClass:"weui-switch",attrs:{type:"checkbox",disabled:n.disabled},domProps:{checked:Array.isArray(n.currentValue)?n._i(n.currentValue,null)>-1:n.currentValue},on:{change:function(e){var t=n.currentValue,i=e.target,a=!!i.checked;if(Array.isArray(t)){var r=n._i(t,null);i.checked?r<0&&(n.currentValue=t.concat([null])):r>-1&&(n.currentValue=t.slice(0,r).concat(t.slice(r+1)))}else n.currentValue=a}}}),n._v(" "),n.preventDefault?t("div",{staticClass:"vux-x-switch-overlay",on:{click:n.onClick}}):n._e()])])},a=[],r={render:i,staticRenderFns:a};e.a=r},743:function(n,e,t){"use strict";var i=t(441),a=t(40),r=t(63);e.a={components:{XSwitch:i.a,Group:a.a,Cell:r.a},methods:{onClick:function(n,e){var t=this;console.log(n,e),this.$vux.loading.show({text:"in processing"}),setTimeout(function(){t.$vux.loading.hide(),t.value2=n},1e3)}},data:function(){return{value1:!0,value2:!1,stringValue:"0"}}}},744:function(n,e,t){"use strict";var i=function(){var n=this,e=n.$createElement,t=n._self._c||e;return t("div",[t("group",{attrs:{title:n.$t("value map")}},[t("x-switch",{attrs:{title:n.$t("default true"),"value-map":["0","1"]},model:{value:n.stringValue,callback:function(e){n.stringValue=e},expression:"stringValue"}}),n._v(" "),t("cell",{attrs:{title:"value",value:typeof n.stringValue+" "+n.stringValue}})],1),n._v(" "),t("group",{attrs:{title:n.$t("Basic Usage")}},[t("x-switch",{attrs:{title:n.$t("default false")}}),n._v(" "),t("x-switch",{attrs:{title:n.$t("default true"),"inline-desc":n.value1+""},model:{value:n.value1,callback:function(e){n.value1=e},expression:"value1"}})],1),n._v(" "),t("group",{attrs:{title:n.$t("disabled")}},[t("x-switch",{attrs:{title:n.$t("default false"),disabled:""}}),n._v(" "),t("x-switch",{attrs:{title:n.$t("default true"),value:!0,disabled:""}})],1),n._v(" "),t("group",{attrs:{title:n.$t("prevent default")}},[t("x-switch",{attrs:{title:n.$t("default false"),"prevent-default":""},on:{"on-click":n.onClick},model:{value:n.value2,callback:function(e){n.value2=e},expression:"value2"}})],1),n._v(" "),t("group",{attrs:{title:n.$t("html title")}},[t("x-switch",{attrs:{disabled:"",title:n.$t("switch red")}})],1)],1)},a=[],r={render:i,staticRenderFns:a};e.a=r}});