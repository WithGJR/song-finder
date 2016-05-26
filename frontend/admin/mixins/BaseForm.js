module.exports = {
  computed: {
    formType(){
      return (this.$route.params.id === undefined) ? 'new' : 'edit';
    },
    title(){
      return (this.formType === 'new') ? this.newFormTitle : this.editFormTitle;
    },
    btnText(){
      return (this.formType === 'new') ? this.newFormBtnText : this.editFormBtnText;
    },
    newFormTitle() {
      return '';
    },
    editFormTitle() {
      return '';
    },
    newFormBtnText() {
      return '';
    },
    editFormBtnText() {
      return '';
    },
    currentId(){
      return parseInt(this.$route.params.id);
    }
  }
};
