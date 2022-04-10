<template>
  <div class="my-4" align="center">
        <b-card
            align="center"
            title="Register"
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
        >
            <b-form @submit="onSubmit">

                <b-card-text>
                    <b-form-group align="left" label="Username">
                        <b-form-input type="text" placeholder="Enter ID" v-model="form.username" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Password">
                        <b-form-input type="password" placeholder="********" v-model="form.password"></b-form-input>
                    </b-form-group>

                    
                    <b-form-group align="left" label="First Name" >
                        <b-form-input type="text" placeholder="สมหมาย" v-model="form.first_name"></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Last Name">
                        <b-form-input type="text" placeholder="ไม่สมหวัง" v-model="form.last_name"></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Age">
                        <b-form-spinbutton id="sb-inline" v-model="form.age" inline></b-form-spinbutton>
                    </b-form-group>

                    <b-form-group align="left" label="Image URL">
                        <b-form-input type="text" placeholder="https://sadlasjdkas.com" v-model="form.image_src"></b-form-input>
                    </b-form-group>
                    
                </b-card-text>
                
                <b-button type="submit" variant="primary">Submit</b-button>
            </b-form>
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
                first_name:"",
                last_name:"",
                age:0,
                image_src:"",
            },
        }
    },
    methods:{
        async onSubmit(evt) {
            evt.preventDefault()
            await axios.post('http://localhost:8000/signup',JSON.stringify({
                username:this.form.username,
                password:this.form.password,
                userdetail: {
                    first_name:this.form.first_name,
                    last_name:this.form.last_name,
                    image_src:this.form.image_src,
                    age:this.form.age,
                }
            }),{
                headers: {'Content-Type': 'application/json'}
            })
            .then((res) => {
                console.log(res)
            })
            .catch((err) => {
                console.log(err)
            })
            // try {
            //     await fetch('http://localhost:8000/signup', {
            //         method:'POST',
            //         body: JSON.stringify({
            //             username:this.form.username,
            //             password:this.form.password,
            //             userdetail: {
            //                 first_name:this.form.first_name,
            //                 last_name:this.form.last_name,
            //                 age:this.form.age,
            //                 image_src:this.form.image_src,
            //                 auth_jwt:"",
            //             }
            //         })
            //     })
            //     // await this.$router.push('/login')
            // } catch (err) {
            //     console.log(err)
            // }
        }
    }
}
</script>