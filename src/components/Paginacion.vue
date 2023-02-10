<template>
    <div class="text-left absolute ml-1 mt-[6px]">
        <label>Total de páginas</label><input id="id_total_pag" type="text" class="bg-zinc-100 ml-2 p-2 focus:outline-none text-sm w-20 h-8 border border-b-black" placeholder="" value="10" >
    </div>
    <div class="inline-flex rounded-md shadow-sm absolute left-0 mt-1 ml-[230px]" role="group">
        <button id="id_first" @click="clickFirst()" type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-l-lg hover:bg-gray-100 hover:text-gray-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white" >
        </button>
        <button id="id_prev" @click="clickPrev()" type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 hover:bg-gray-100 hover:text-gray-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
        </button>
        <button id="id_next" @click="clickNext()" type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-r-md hover:bg-gray-100 hover:text-gray-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
        </button>
        <button id="id_last" @click="clickLast()" type="button" class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-r-md hover:bg-gray-100 hover:text-gray-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white">
        </button>
    </div>
    <div class="text-left absolute ml-1 mt-[48px]">
        <label>Registros máximos a devolver</label><input id="id_max_num" type="text" class="bg-zinc-100 ml-2 p-2 focus:outline-none text-sm w-20 h-8 border border-b-black" placeholder="" value="40" >
    </div>
    <!-- <paginate
    v-model="page"
    :page-count="5"
    :page-range="3"
    :margin-pages="2"
    :click-handler="clickCallback"
    :first-last-button="true"
    :first-button-text="'<<'"
    :last-button-text="'>>'"
    :prev-text="'<'"
    :next-text="'>'"
    :container-class="'pagination'"
    :page-class="'page-item'"
    class=" cursor-pointer absolute py-[6px] ml-[324px]"
    >
  </paginate> -->
</template>
<script setup>
    defineProps({
        // max_num:"",
        itemsTotales:[],
        pageNum:1
    })
</script>
<script>
// import Paginate from "vuejs-paginate-next";
export default {
  name: "Tabla",
  data(){
    return {
        pageNumReturn:1,
        total_pag:""
    };
  },
//   components: {
//       paginate: Paginate,
//     },
  mounted() {
    document.getElementById("id_first").innerHTML="|<"
    document.getElementById("id_prev").innerHTML="<"
    document.getElementById("id_next").innerHTML=">"
    document.getElementById("id_last").innerHTML=">|"
    const id_total_pag = document.getElementById("id_total_pag")
    this.total_pag = id_total_pag.value
    this.$emit('total_pag', this.total_pag)
    id_total_pag.addEventListener('input', 
        async (e)=>{this.$emit('total_pag',e.target.value)}
    );
    const id_max_num = document.getElementById("id_max_num")
    this.$emit('max_num', id_max_num.value)
    id_max_num.addEventListener('input', 
        async (e)=>{this.$emit('max_num',e.target.value)}
    );
  },
  methods: {
    clickPrev(){
        if(this.pageNum>1 && this.itemsTotales.length!=0){
            this.pageNumReturn = this.pageNum - 1
            this.$emit('pageNum',this.pageNumReturn)
        }
        
    },
    clickNext(){
        const id_total_pag = document.getElementById("id_total_pag")
        this.total_pag = id_total_pag.value
        let t = parseInt(this.total_pag)
        console.log(this.total_pag)
        if(this.pageNum < t && this.itemsTotales.length!=0){
            this.pageNumReturn = this.pageNum + 1
            this.$emit('pageNum',this.pageNumReturn)
        }

    },
    clickLast(){
        const id_total_pag = document.getElementById("id_total_pag")
        this.total_pag = id_total_pag.value
        let t = parseInt(this.total_pag)
        if(this.pageNum < t && this.itemsTotales.length!=0){
            this.pageNumReturn = t
            this.$emit('pageNum',this.pageNumReturn)
        }
    },
    clickFirst(){
        if(this.pageNum>1 && this.itemsTotales.length!=0){
            this.pageNumReturn = 1
            this.$emit('pageNum',this.pageNumReturn)
        }
    },
  }
}
</script>
<style lang="css">
  @import "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css";
</style>