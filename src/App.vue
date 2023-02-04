<script setup>
import Tabla from './components/Tabla.vue'
import Buscador from './components/Buscador.vue'
import Titulo from './components/Titulo.vue'
import Contenido from './components/Contenido.vue'
</script>
<template>
      <div class="Encabezado fixed header border border-gray-600 bg-gray-600">
      <Titulo/>
    </div>
    <div class="Buscador fixed header">
      <Buscador @Items="handleBuscador" />
    </div>
  <div class="fixed main grid grid-cols-12 gap-2 bg-zinc-50" >
    <div class="overflow-y-scroll col-span-8 rounded-lg border border-gray-500 bg-gray-200 p-8 sm:col-span-8">
      <Tabla v-bind:items="Items" @Msg="handleTabla"/>
    </div>
    <div class="overflow-y-scroll col-span-4 rounded-lg border border-gray-500 bg-gray-100 p-8 sm:col-span-4">
      <Contenido v-bind:msg="Msg" />
    </div>
  </div>
  <div class="fixed footer-copyright footer border-gray-600 bg-gray-600 ">
    <br><p class="text-zinc-50"> © 2023 Creado por Ivan Rodriguez</p>
  </div>
  </template>
<script>
export default {
  name: 'App',
  data() {
    return {
      Msg: '',
      Items:[]
    }
  },
  methods: {
    handleBuscador (value) {
      this.Items = value
    },
    handleTabla (value) {
      this.Msg = value
    }
  }
}
</script>
<style scoped>
  .main{
    left: 3%;
    top:18%;
    width: 94%;
    height: 70%;
  }
  .Encabezado{
    left: 0;
    top: 0;
    width: 100%;
    height:10%;
  }
  .Buscador{
    left: 0;
    top: 10%;
    width: 100%;
    height:10%;
  }
  .footer-copyright{
    left: 0;
    top: 90%;
    width: 100%;
    height: 10%;
  }
</style>
