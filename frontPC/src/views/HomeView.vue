<template >
  <div class="surface-card p-4 shadow-2 border-round w-full  container-login">
    <div class="container-content-login">
      <div class="text-center mb-5">
      </div>
      <div>
        <label for="name1" class="block text-900 font-medium mb-2">Name</label>
        <pv-input-text :id="name1"  type="text" class="w-full mb-3" />


        <label for="lastname" class="block text-900 font-medium mb-2">LastName</label>
        <pv-input-text :id="lastname" type="text" class="w-full mb-3" />

        <label for="cci" class="block text-900 font-medium mb-2">CCI</label>
        <pv-input-text :id="cci"  type="text" class="w-full mb-3" />


        <label for="age" class="block text-900 font-medium mb-2">Age</label>
        <pv-input-text :id="age" type="text" class="w-full mb-3" />

        <pv-button label="Registrar nuevo Propietario" icon="pi pi-user" class="w-full" @click="logeado()"></pv-button>
      </div>
    </div>
  </div>
</template>

<script>


export default {
  name: "HomeView.component",
  let socket=new WebSocket("ws://localhost:8080/ws")

  created() {

  },
  data() {
    return {
      name1: "",
      lastname: "",
      cci:"",
      age:""
    }
  },
  methods:{
    logeado(){

      let data={
        "name":this.name1,
        "password":this.lastname,
        "CCI":this.cci,
        "age":this.age
      };


    },


  }
};
</script>

<style scoped>
.container1
{
  position: center;
  bottom: 0px;
  padding-top: 10px;
  padding-left: 650px;
}
.p-float-label{
  position: center;
  bottom: 0px;
  padding-top: 150px;
  padding-left: 650px;
}
.container-login{
  height: 100vh;
  width: 100%;
  display: flex;
  justify-content:center;
}
.container-content-login{
  width: 500px;
}
</style>
