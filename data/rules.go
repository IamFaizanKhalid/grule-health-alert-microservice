package data

// rules definitions

var Drls = `
rule updateBMI "Update BMI for every user" salience 10 {
    when
		PERSON.Bmi != PERSON.Weight/ ((PERSON.Height/100)*(PERSON.Height/100))
    then
		Log("BMI of " + PERSON.Name + " updated to " + PERSON.UpdateBMI());
        Retract("updateBMI");
}
`
