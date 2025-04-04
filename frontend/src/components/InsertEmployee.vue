<template>
    <div class="container">
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label>Employee Name</label>
                <input type="text" v-model="localName" required placeholder="Enter name of employee">
            </div>
            <div class="form-group">
                <label>Age</label>
                <input type="number" v-model="localAge" required placeholder="Enter your age">
            </div>
            <div class="form-group">
                <label>Position</label>
                <input type="text" v-model="localPosition" required placeholder="Enter your company position">
            </div>
            <div class="form-group">
                <label>Salary</label>
                <input type="number" v-model="localSalary" placeholder="Rs. 0">
            </div>
            <button>{{ isUpdate ? "Update Employee" : "Upload Data" }}</button>
        </form>
    </div>
</template>

<script setup>
import { ref, defineProps, defineEmits, watch } from "vue";
import { AddEmployee, UpdateEmployee } from "../../wailsjs/go/backend/EmployeeService";
import { useRouter } from "vue-router";
import { useToast } from 'vue-toastification'

const router = useRouter();
const toast = useToast();

const props = defineProps({
    id: Number, // Employee ID (only for update)
    name: String,
    age: Number,
    position: String,
    salary: Number,
    isUpdate: Boolean, // To differentiate between Add & Update mode
});

const emit = defineEmits(["employeeUpdated"]); // Emit event after update

// Create local refs to handle form data separately
const localName = ref(props.name || "");
const localAge = ref(props.age || "");
const localPosition = ref(props.position || "");
const localSalary = ref(props.salary || "");

// Watch for changes in props and update local state
watch(
  () => [props.name, props.age, props.position, props.salary],
  ([newName, newAge, newPosition, newSalary]) => {
    localName.value = newName || "";
    localAge.value = newAge || "";
    localPosition.value = newPosition || "";
    localSalary.value = newSalary || "";
  },
  { immediate: true }
);


const submitForm = async () => {
    const salaryValue = parseInt(localSalary.value) || 0;
    try {
        if (props.isUpdate) {
            // ✅ Update existing employee
            await UpdateEmployee(
                props.id,
                localName.value,
                parseInt(localAge.value),
                localPosition.value,
                salaryValue
            );
            toast.info("Employee data is updated!")
            
        } else {
            // ✅ Add new employee
            await AddEmployee(
                localName.value,
                parseInt(localAge.value),
                localPosition.value,
                parseInt(localSalary.value)
            );
            toast.success("Employee added successfully!")
            router.push("/"); // Redirect back to employee list
        }

        // Emit event for update
        emit("employeeUpdated");

        // Reset form (only if adding)
        if (!props.isUpdate) {
            localName.value = "";
            localAge.value = "";
            localPosition.value = "";
            localSalary.value = "";
        }
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