<template>

</template>

<script>
import axios from 'axios'


export default {
  name: "ques1",
  mounted: function(){
    axios.get("https://dummyjson.com/products/1").then((response) => {
      console.log("response is ::: ",response.data.title);
    });
  }
}
</script>

<style scoped>

</style>