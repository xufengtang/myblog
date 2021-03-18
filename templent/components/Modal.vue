<style lang="scss">
.modal {
 &__mask {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
 }
 &__main {
  margin: 0 auto;
  margin-bottom: 5%;
  margin-top: 20%;
  width: 500px;
  background: #fff;
  border-radius: 8px;
 }
 /* 省略部分样式 */
}
</style>
<template>
 <div class="modal__mask">
  <div class="modal__main">
   <div class="modal__header">
    <h3 class="modal__title">弹窗标题</h3>
    <span class="modal__close">x</span>
   </div>
   <div class="modal__content">
    弹窗文本内容
   </div>
   <div class="modal__footer">
    <button>取消</button>
    <button>确认</button>
   </div>
  </div>
 </div>
</template>
 
<script>
export default {
 setup() {
  return {};
 },
};
</script>