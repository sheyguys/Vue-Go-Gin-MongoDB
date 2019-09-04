<template>
    <b-form class="form">
        <center>
        <h1>Vue Gin MongoDB</h1>
        <img src="../assets/logo.png"><br>
        <img src="../assets/gin.png"><br>
        <img src="../assets/mongodb.png"><br>
        <br><br>
        <b-button @click="$router.push('register')" variant="warning">Register</b-button>
        <br><br>
        <b-button @click="$router.push('result')" variant="success">Result</b-button>
        </center>
    </b-form>
</template>
<script>

export default {
    
}
</script>>
<style>
h1 {
  text-align: center;
}

.form{
  margin-left: 10px;
  margin-right: 10px;
  margin-top: 10px;
  background-color: #000000;
  color: white;
  padding: 40px;
  font-family: Courier;
}

</style>