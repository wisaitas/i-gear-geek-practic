<template>
  <div class="my-3" align="center">
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
            <!-- <b-card header="result" class="mt-3">password : {{this.form.password}}</b-card> -->
        </b-card>
    </div>
</template>

<script>
export default {
    data(){
        return {
            form:{
                username:"",
                password:"",
            },
        }
    },
    methods:{
        async onSubmit(evt){
            evt.preventDefault()
            try {
                const response = await this.$auth.loginWith('local',{data: this.form})
                await this.$router.push('/')
                console.log(response)
            } catch (err) {
                console.log(err)
            }
        }
    }
}
</script>