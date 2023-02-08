<template>
    <div>
        <input id="id_max_num" type="text" class="absolute bg-zinc-100 p-1 focus:outline-none left-3 mt-2 text-sm w-32 h-8 border border-b-black" placeholder="Registros máximos" value="20" >
    </div>
</template>
<script>
export default {
  name: "Tabla",
  mounted() {
    var id_max_num = document.getElementById("id_max_num")
    this.$emit('max_num', id_max_num.value)
    id_max_num.addEventListener('input', 
        async (e)=>{this.$emit('max_num',e.target.value)}
    );
  }

};
</script>