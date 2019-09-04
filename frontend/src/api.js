import axios from 'axios'

const SERVER_URL = 'http://localhost:8081';

export default {
    postMember(member_name_th, member_name_eng, member_idcard, member_facebook, member_email, member_address, member_phone) {
        const url = SERVER_URL + `/employee`;
        return axios.post(url, {
            member_name_th:  member_name_th,
            member_name_eng: member_name_eng,
            member_idcard:  member_idcard,
            member_facebook:  member_facebook,
            member_email: member_email,
            member_address:  member_address,
            member_phone:  member_phone,
        }).then(function (response) {
            console.log(response.data)
        })
            .catch(function (error) {
                console.log(error)
            });
    },

    getallmember() {
        const url = SERVER_URL + `/employee`;
        return axios.get(url).then(response => response.data);
    },

    deletemember(id){
        const url = SERVER_URL + `/employee/` + id;
        return axios.delete(url).then(function (response) {
            console.log("Deleted"+response)
        })
            .catch(function (error) {
                console.log(error)
            });
    }, 
}