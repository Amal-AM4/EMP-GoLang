export namespace backend {
	
	export class Employee {
	    id: number;
	    name: string;
	    age: number;
	    position: string;
	    salary: number;
	
	    static createFrom(source: any = {}) {
	        return new Employee(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.age = source["age"];
	        this.position = source["position"];
	        this.salary = source["salary"];
	    }
	}

}

