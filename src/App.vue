<script setup>
import Tabla from './components/Tabla.vue'
import Buscador from './components/Buscador.vue'
import Titulo from './components/Titulo.vue'
import Contenido from './components/Contenido.vue'
import Paginacion from './components/Paginacion.vue';
</script>
<template>
    <div class="left-0 top-0 w-full h-[95px] fixed header border border-gray-600 bg-gray-600">
      <Titulo/>
    </div>
    <div class="left-0 top-[100px] w-full h-[10%] fixed header">
      <Buscador @Items="handleBuscador" />
    </div>
  <div class="fixed left-[1%] top-[178px] w-[98%] h-[70%] rounded-lg grid grid-cols-12 gap-2" >
    <div class=" col-span-8 rounded-lg border border-gray-500 p-2 sm:col-span-8  bg-gray-200">
      <Tabla v-bind:items="Items" @Msg="handleTabla"/>
      <Paginacion />
    </div>
    <div class="overflow-y-scroll col-span-4 rounded-lg border border-gray-500 bg-gray-100 p-6 sm:col-span-4">
      <Contenido v-bind:msg="Msg" />
    </div>
  </div>
  <div class="fixed left-0 top-[91%] w-full h-[9%] footer border-gray-600 bg-gray-600 ">
    <br><p class="text-zinc-50"> © 2023 Creado por Ivan Rodriguez</p>
  </div>
  </template>
<script>
export default {
    name: "App",
    data() {
        return {
            Msg: "",
            Items: []
        };
    },
    methods: {
        handleBuscador(value) {
            this.Items = value;
            this.Msg = "";
        },
        handleTabla(value) {
            this.Msg = value;
        }
    },
    components: { Paginacion }
}
</script>
<style scoped>
</style>
