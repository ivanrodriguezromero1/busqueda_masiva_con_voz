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
      <Buscador v-bind:max_num="Max_num" @ItemsTotales="handleBuscador" />
    </div>
  <div class="fixed left-[1%] top-[178px] w-[98%] h-[69%] rounded-lg grid grid-cols-12 gap-2" >
    <div class=" col-span-8 rounded-lg border border-gray-500 p-2 sm:col-span-8  bg-gray-200">
      <Tabla v-bind:items="Items" @Msg="handleTabla"/>
      <Paginacion v-bind:pageNum="PageNum" v-bind:itemsTotales="ItemsTotales" @max_num="handleMaxNum" @total_pag="handleTotalPag" @pageNum="handlePageNum" />
    </div>
    <div class="overflow-y-scroll col-span-4 rounded-lg border border-gray-500 bg-gray-100 p-6 sm:col-span-4">
      <Contenido v-bind:msg="Msg" />
    </div>
  </div>
  <div class="fixed left-0 top-[93%] w-full h-[7%] footer border-gray-600 bg-gray-600 ">
    <br><p class="text-zinc-50 mt-[-10px]"> © 2023 Creado por Ivan Rodriguez</p>
  </div>
  </template>
<script>
export default {
    name: "App",
    data() {
        return {
          Max_num: "",
          ItemsTotales: [],
          Items: [],
          PageNum: 1,
          Total_pag:"",
          TotalPag:10,
          Msg: "" 
        };
    },
    methods: {
        handleBuscador(ItemsTotales) {
          this.ItemsTotales = ItemsTotales
          this.PageNum = 1
          this.Items = this.Filtro(ItemsTotales,this.PageNum)
          // console.log(this.PageNum)
          this.Msg = "";
        },
        handleMaxNum(Max_num){
          this.Max_num = Max_num;
        },
        handleTotalPag(Total_pag){
          this.Total_pag = Total_pag;
        },
        handleTabla(Msg) {
          this.Msg = Msg;
        },
        handlePageNum(pageNumReturn){
          this.PageNum = pageNumReturn
          this.Items = this.Filtro(this.ItemsTotales,pageNumReturn);
          // console.log("=========")
          // console.log(this.Items)
        },
        Filtro(ItemsTotales,pageNumReturn){
          // let numXPag = (this.Max_num%totalPag==0) ? (this.Max_num/totalPag):((this.Max_num/totalPag)+1)
          let m = parseInt(this.Max_num)
          let t = parseInt(this.Total_pag)
          console.log(Math.ceil(m/t))
          let numXPag = Math.ceil(m/t) 
          let tope = (pageNumReturn != t) ? (numXPag*pageNumReturn) : m
          // console.log("=========")
          // console.log(this.Max_num)
          // console.log(PageNum)
          // console.log(numXPag)
          // console.log(tope)
          return ItemsTotales.slice(numXPag*(pageNumReturn-1),tope)
        }
    },
}
</script>
<style scoped>
</style>
