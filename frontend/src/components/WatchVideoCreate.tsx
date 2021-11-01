import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { CaseInterface } from "../models/ICase";
import { AssessInterface } from "../models/IAssess";
import { SymptomInterface } from "../models/ISymptom";
import { StateInterface } from "../models/IState";
import { AssessmentsheetInterface} from "../models/IAssessmentsheet";


import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function AssessmentsheetCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());

  const [cases, setCase] = useState<CaseInterface[]>([]);
  const [assess, setAssess] = useState<AssessInterface[]>([]);
  const [symptom, setSymptom] = useState<SymptomInterface[]>([]);
  const [state, setState] = useState<StateInterface[]>([]);
  const [ams, setAssessmentsheet] = useState<Partial<AssessmentsheetInterface>>(
    {}
  );
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  /*const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof ams;
    const { value } = event.target;
    setAssessmentsheet({ ...ams, [id]: value });
  };*/

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof ams;
    setAssessmentsheet({
      ...ams,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

 

  const getCase = async () => {
    fetch(`${apiUrl}/cases`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCase(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getAssess = async () => {
    fetch(`${apiUrl}/assesses`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setAssess(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getState = async () => {
    fetch(`${apiUrl}/states`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setState(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getSymptom = async () => {
    fetch(`${apiUrl}/symptoms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSymptom(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {

    getCase();
    getAssess();
    getState();
    getSymptom();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
    CaseID:  convertType(ams.CaseID),
     AssessID:  convertType(ams.AssessID),
     StateID:   convertType(ams.StateID),
     SymptomID:   convertType(ams.SymptomID),
     Assesstime:  selectedDate, 
    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/assessment_sheets`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
        }
      });
  }

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกใบประเมินอาการ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่อ</p>
              <Select
                native
                value={ams.CaseID}
                onChange={handleChange}
                inputProps={{
                  name: "CaseID",
                }}
              >
                <option aria-label="None" value="">
                  ชื่อผู้ป่วย
                </option>
                {cases.map((item: CaseInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>อาการ</p>
              <Select
                native
                value={ams.SymptomID}
                onChange={handleChange}
                inputProps={{
                  name: "SymptomID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกอาการ
                </option>
                {symptom.map((item: SymptomInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.SymptomData}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>สถานะ</p>
              <Select
                native
                value={ams.StateID}
                onChange={handleChange}
                inputProps={{
                  name: "StateID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกสถานะ
                </option>
                {state.map((item: StateInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.StateData}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเมินอาการ</p>
              <Select
                native
                value={ams.AssessID}
                onChange={handleChange}
                inputProps={{
                  name: "AssessID",
                }}
              >
                <option aria-label="None" value="">
                  เลือกอาการ
                </option>
                {assess.map((item: AssessInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.AssessData}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="AssessTime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/watch_videos"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default AssessmentsheetCreate;
