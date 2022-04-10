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
                        <b-form-input type="text" placeholder="https://zxczjclzxcza.com" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="First Name">
                        <b-form-input type="text" placeholder="สมปอง" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Last Name">
                        <b-form-input type="text" placeholder="ห้าวจัด" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Age">
                        <b-form-spinbutton id="sb-inline" inline></b-form-spinbutton>
                    </b-form-group>

                </b-card-text>
                
                <b-button @click="onSubmit" variant="primary">Save</b-button>
            </b-form>
            </b-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            details:{
                first_name:"",
                last_name:"",
                image:"",
                age:0,
            }
        }
    },
    layout: "navbar",
    middleware: 'auth',
    methods:{
        onSubmit() {
            this.$router.push("/profile")
        }
    }
}
</script>