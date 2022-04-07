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
                        <b-form-input type="text" placeholder="Enter ID" ></b-form-input>
                    </b-form-group>

                    <b-form-group align="left" label="Password">
                        <b-form-input type="password" placeholder="********" ></b-form-input>
                    </b-form-group>
                </b-card-text>

                <b-button type="submit" variant="primary">Login</b-button>
            </b-form>
        </b-card>
    </div>
</template>

<script>
export default {
    methods:{
        onSubmit(evt){
            evt.preventDefault()
            console.log("hello world")
        }
    }
}
</script>