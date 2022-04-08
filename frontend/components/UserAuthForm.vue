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
            <b-card header="result" class="mt-3">token : {{this.token}}</b-card>
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
            token: "",
        }
    },
    methods:{
        // async handlerLoginClicked() {
        //     try {
        //         const response = await this.$auth.loginWith('local',{data: {username: this.form.username,password: this.form.password}})
        //         console.log(response)
        //     } catch (error) {
        //         console.log(error)
        //     }
        // }
        async onSubmit(evt){
            try {
                const response = await this.$auth.loginWith('local',{data: {username:this.fomr.username},passowrd:this.form.password})
                console.log(err)
            } catch (err) {
                console.log(err)
            }
            // evt.preventDefault()
            // const res = await axios.post('http://localhost:8000/login',{
            //     username:this.form.username,
            //     password:this.form.password
            // })
            // .then(res => {
            //     localStorage.setItem('token',res.data.jwtToken)
            //     this.token = localStorage.getItem('token')
            // })
            // .catch(err => {
            //     console.log(err)
            // })
        }
    }
}
</script>