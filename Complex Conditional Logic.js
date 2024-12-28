let person = { age: 25, experience: 3, skills: ["JavaScript", "Python"] };
if (person.age >= 18 && person.experience >= 2 && person.skills.includes("JavaScript")) {
    console.log("Eligible");
} else {
    console.log("Not Eligible");
}
