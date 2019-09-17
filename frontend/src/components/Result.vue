<template>
      <b-form class="form">
      <h1>EMPLOYEE RECORDS</h1>
      <br>
      <table border="1" align="center" width="80%" cellspacing="0">
          <tr>
            <th>Eng Name</th>
            <th>Thai Name</th>
            <th>ID.card</th>
            <th>Phone</th>
            <th>Address</th>
            <th>Email</th>
            <th>Facebook</th>
            <th>DELETE</th>
          </tr>
          <tr v-for="(employee) in employees" :key="employee.member_id">
            <td>{{employee.member_name_eng}}</td>
            <td>{{employee.member_name_th}}</td>
            <td>{{employee.member_idcard}}</td>
            <td>{{employee.member_phone}}</td>
            <td>{{employee.member_address}}</td>
            <td>{{employee.member_email}}</td>
            <td>{{employee.member_facebook}}</td>
            <th><b-button @click="Delete(employee.member_id)" variant="danger">DELETE</b-button></th>
          </tr>
      </table>

      </b-form>
</template>
<script>
import api from "../api";
export default {
    data() {
    return {
      employees: []
    };
  },
  methods: {
    getallmember() {
      api.getallmember().then(result => {
        console.log(result.member);
        this.employees = result.member;
      });
    },

    Delete(id){
      api.deletemember(id)  
        console.log(id);
        window.location.reload();
        alert("Delete Complete!!");
    }
  },
  mounted() {
    this.getallmember();
  }
}
</script>
<style>
h1 {
  text-align: center;
}

.form{
  margin-left: 10px;
  margin-right: 10px;
  margin-top: 10px;
  background-color: #282c34;
  color: white;
  padding: 40px;
  font-family: Courier;
}

</style>