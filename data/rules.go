package data

// rules definitions

var Drls = `
rule updateBMI "Update BMI for every user" salience 10 {
    when
		PERSON.Bmi != PERSON.Weight/ ((PERSON.Height/100)*(PERSON.Height/100))
    then
		PERSON.Bmi = PERSON.Weight/ ((PERSON.Height/100)*(PERSON.Height/100));
		Log("BMI of " + PERSON.Name + " updated to " + PERSON.Bmi);
        Retract("updateBMI");
}
`
