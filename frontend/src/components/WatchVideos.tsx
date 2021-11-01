import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
//import { WatchVideoInterface } from "../models/IWatchVideo";
import { format } from 'date-fns'
import { AssessmentsheetInterface } from "../models/IAssessmentsheet";


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Assessmentsheet() {
  const classes = useStyles();
  const [ams, setAssessmentsheet] = useState<AssessmentsheetInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getAssessmentsheet = async () => {
    fetch(`${apiUrl}/assessment_sheets`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setAssessmentsheet(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getAssessmentsheet();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการประเมินอาการ
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/watch_video/create"
              variant="contained"
              color="primary"
            >
              สร้างใบประเมินอาการ
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="20%">
                  ชื่อ
                </TableCell>
                <TableCell align="center" width="20%">
                  อาการ
                </TableCell>
                <TableCell align="center" width="15%">
                  สถานะ
                </TableCell>
                <TableCell align="center" width="20%">
                  ประเมิน
                </TableCell>
                <TableCell align="center" width="20%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {ams.map((item: AssessmentsheetInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Case.Name}</TableCell>
                  <TableCell align="center">{item.Symptom.SymptomData}</TableCell>
                  <TableCell align="center">{item.State.StateData}</TableCell>
                  <TableCell align="center">{item.Assess.AssessData}</TableCell>
                  <TableCell align="center">{format((new Date(item.AssessTime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Assessmentsheet;
