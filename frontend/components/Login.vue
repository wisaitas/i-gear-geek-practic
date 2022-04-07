<template>
  <div class="ml-3 my-3">
        <b-card
            align="center"
            title="Login"
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
        >
            <b-form @submit="onSubmit">

                <b-card-text>
                    <b-form-group align="left" label="ID">
                        <b-form-input type="text" placeholder="Enter ID" v-model="form.username" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Password">
                        <b-form-input type="password" placeholder="********" v-model="form.password"></b-form-input>
                    </b-form-group>
                </b-card-text>

                <b-button type="submit" variant="primary">Login</b-button>
            </b-form>
            <b-card header="result" class="mt-3">token : {{token}}</b-card>
        </b-card>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    data(){
        return {
            form:{
                username:"",
                password:"",
            },
            token:"",
        }
    },
    methods:{
        async onSubmit(evt){
            evt.preventDefault()
            const json = JSON.stringify({username:this.form.username,password:this.form.password})
            console.log(json)
            const res = await axios.post('localhost:8000/login',json,{
                headers: {
                    // Overwrite Axios's automatically set Content-Type
                    'Content-Type': 'application/json'
                }
            })
            .then(function (res) {
                // this.token
                console.log(res)
            })
            .catch(function (err) {
                console.log(err)
            })
        }
    }
}
</script>