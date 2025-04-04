<template>
    <div class="employee-container">
        <h2>Employee List</h2>
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Age</th>
                    <th>Position</th>
                    <th>Salary</th>
                    <th>Edit</th>
                    <th>Remove</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(employee, index) in employees" :key="employee.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ employee.name }}</td>
                    <td>{{ employee.age }}</td>
                    <td>{{ employee.position }}</td>
                    <td>₹{{ employee.salary.toLocaleString() }}</td>
                    <td>
                        <RouterLink :to="`/edit/${employee.id}`" class="edit-btn">Edit</RouterLink>
                    </td>
                    <td>
                        <button @click="removeEmployee(employee.id)" class="remove-btn">Remove</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup>
import { RouterLink } from 'vue-router'
import { onMounted, ref } from "vue";
import { GetEmployees, DeleteEmployee } from '../../wailsjs/go/backend/EmployeeService'

const employees = ref([])

const fetchEmployees = async () => {
    try {
        employees.value = await GetEmployees();
    } catch (error) {
        console.error("Error fetching employees:", error);
    }
};

const removeEmployee = async (id) => {
    if (confirm("Are you sure you want to remove this employee?")) {
        try {
            await DeleteEmployee(id); // ✅ Call Wails API
            employees.value = employees.value.filter(emp => emp.id !== id); // ✅ Update UI
        } catch (error) {
            console.error("Failed to delete employee:", error);
        }
    }
};

onMounted(fetchEmployees);
</script>

<style lang="scss" scoped>
.employee-container {
    max-width: 90%;
    margin: 20px auto;
    text-align: center;

    h2 {
        margin-bottom: 20px;
        font-size: 24px;
        color: #2E6049;
    }

    table {
        width: 100%;
        border-collapse: collapse;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
        background: #fff;
        border-radius: 8px;
        overflow: hidden;

        thead {
            background: #2E6049;
            color: #fff;
            text-transform: uppercase;
        }

        th, td {
            padding: 12px;
            border: 1px solid #ddd;
            text-align: center;
            text-transform: capitalize;
        }

        tbody tr:nth-child(even) {
            background: #f9f9f9;
        }

        tbody tr:hover {
            background: #e8f5e9;
        }
    }

    .edit-btn {
        text-decoration: none;
        color: #fff;
        background: #0288D1;
        padding: 6px 12px;
        border-radius: 4px;
        transition: 0.3s;

        &:hover {
            background: #0277BD;
        }
    }

    .remove-btn {
        border: none;
        padding: 6px 12px;
        background: #D32F2F;
        color: white;
        border-radius: 4px;
        cursor: pointer;
        transition: 0.3s;

        &:hover {
            background: #B71C1C;
        }
    }
}
</style>
