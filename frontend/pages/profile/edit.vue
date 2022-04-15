<template>
    <div class="my-3" align="center">
        <b-card
            align="center"
            title="Edit Profile"
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
        >
            <b-img rounded="circle" :src="this.$auth.user.image_src" width=100></b-img>

            <b-form>

                <b-card-text>
                    <b-form-group align="left" label="Image URL">
                        <b-form-input type="text" placeholder="https://zxczjclzxcza.com" v-model="details.image_src"></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="First Name">
                        <b-form-input type="text" placeholder="สมปอง" v-model="details.first_name" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Last Name" >
                        <b-form-input type="text" placeholder="ห้าวจัด" v-model="details.last_name"></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Age" >
                        <b-form-spinbutton id="sb-inline" v-model="details.age" inline ></b-form-spinbutton>
                    </b-form-group>

                </b-card-text>
                
                <b-button @click="onSubmit" variant="primary">Save</b-button>
                <!-- <b-card header="result" class="mt-3">{{ JSON.stringify ({
                old_first_name:this.$auth.user.first_name,
                first_name:this.details.first_name,
                last_name:this.details.last_name,
                image_src:this.details.image_src,
                age:this.details.age
            })}}</b-card> -->
            </b-form>
            </b-card>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            details:{
                first_name:"",
                last_name:"",
                image_src:"",
                age:0,
            }
        }
    },
    layout: "navbar",
    middleware: 'auth',
    methods:{
        async onSubmit(evt) {
            evt.preventDefault()
            await axios.put('http://localhost:8000/profile', JSON.stringify ({
                old_first_name:this.$auth.user.first_name,
                first_name:this.details.first_name,
                last_name:this.details.last_name,
                image_src:this.details.image_src,
                age:this.details.age
            }), {
                headers: {'Content-Type': 'application/json'}
            })
            .then((res) => {
                this.$auth.user.first_name = this.details.first_name
                this.$router.push("/profile")
                console.log(res)
            })
            .catch((err) => {
                console.log(err)
            })
        }
    }
}
</script>