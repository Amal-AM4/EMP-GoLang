<template>
    <div class="container">
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label>Employee Name</label>
                <input type="text" v-model="name" required placeholder="Enter name of employee">
            </div>
            <div class="form-group">
                <label>Age</label>
                <input type="number" v-model="age" required placeholder="Enter your age">
            </div>
            <div class="form-group">
                <label>Position</label>
                <input type="text" v-model="position" required placeholder="Enter your company position">
            </div>
            <div class="form-group">
                <label>Salary</label>
                <input type="number" v-model="salary" placeholder="Rs. 0">
            </div>
            <button>Upload Data</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { AddEmployee } from "../../wailsjs/go/backend/EmployeeService";

const name = ref("")
const age = ref("")
const position = ref("")
const salary = ref("")

const submitForm = async () => {
    try {
        const response = await AddEmployee(
            name.value,
            parseInt(age.value),
            position.value,
            parseInt(salary.value)
        );

        name.value = ""
        age.value = ""
        position.value = ""
        salary.value = ""

        alert("Data is inserted"); // ✅ Show success message
    } catch (error) {
        alert("Error: " + error);
    }
};


</script>

<style lang="scss">
.container {
    width: 95%;
    margin: 0 auto;

    form {
        .form-group {
            display: flex;
            flex-direction: column;
            margin-bottom: 16px;

            label {
                color: #111111b2;
                padding-bottom: 4px;
            }

            input {
                padding: 8px 14px;
            }
        }

        button {
            padding: 14px 22px;
            border: none;
            cursor: pointer;
            border-radius: 4px;
            background-color: #2e6049;
            color: #fff;
            transition: all .3s;

            &:hover {
                background-color: #458466;
            }
        }
    }
}
</style>