import { useEffect, useState } from "react";

import dynamic from "next/dynamic";

import "@fullcalendar/core/main.css";
import "@fullcalendar/daygrid/main.css";
import "@fullcalendar/timegrid/main.css";

import { makeStyles } from "@material-ui/core/styles";
import LinearProgress from "@material-ui/core/LinearProgress";

import "./styles.css";

const useStyles = makeStyles(theme => ({
  root: {
    width: "100%",
    "& > * + *": {
      marginTop: theme.spacing(2)
    }
  }
}));

const API = process.env.API;

let Calendar;

export default ({ email, events }) => {
  const [calendarLoaded, setCalendarLoaded] = useState(false);
  const [view, setView] = useState("timeGridWeek");

  // Fullcalendar doesn't support SSR :(
  useEffect(() => {
    Calendar = dynamic({
      modules: () => ({
        calendar: import("@fullcalendar/react"),
        dayGridPlugin: import("@fullcalendar/daygrid"),
        timeGridPlugin: import("@fullcalendar/timegrid")
      }),
      render: (props, options) => {
        const { calendar: Calendar, ...plugins } = options;

        return <Calendar {...props} plugins={Object.values(plugins)} />;
      },
      ssr: false
    });
    setCalendarLoaded(true);
  });

  let CalendarContainer = props => {
    if (!calendarLoaded) {
      return <LinearProgress />;
    }

    return (
      <Calendar
        defaultView={view}
        header={{
          left: "",
          center: "",
          right: "dayGridMonth,timeGridWeek,timeGridDay"
        }}
        datesRender={({ view }) => setView(view.type)}
        events={events}
      />
    );
  };

  return <CalendarContainer />;
};
