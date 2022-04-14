import React, { FunctionComponent, useEffect, useRef, useState } from "react";
import ReactWebcam from "react-webcam";
import styles from "./webcam.module.css";

interface ProcessedImage {
  path: string;
  createdAt: string;
}

const Webcam: FunctionComponent = () => {
  const ref = useRef<ReactWebcam>(null);
  const [screenshots, setScreenshots] = useState<ProcessedImage[]>([]);

  useEffect(() => {
    document.addEventListener("keydown", takeCapture);
    document.addEventListener("keydown", clearImages);

    return () => {
      document.removeEventListener("keydown", takeCapture);
      document.removeEventListener("keydown", clearImages);
    };
  }, []);

  const clearImages = (event: KeyboardEvent) => {
    if (event.code !== "Escape") return;
    setScreenshots([]);
  };

  const takeCapture = async (event: KeyboardEvent) => {
    if (event.code !== "Space") return;
    const imageSrc = ref.current?.getScreenshot();
    if (!imageSrc) return;

    const res: Response = await fetch(imageSrc);
    const blob: Blob = await res.blob();
    const image = await sendImage(blob);
    setScreenshots((prev) => {
      const copy = [...prev];
      copy.push(image);
      return copy;
    });
  };

  const sendImage = async (image: Blob): Promise<ProcessedImage> => {
    const formData = new FormData();
    formData.append("image", image);
    const res = await fetch("http://localhost:8000/", {
      method: "POST",
      body: formData,
    });
    return await res.json();
  };

  return (
    <div className={styles["container"]}>
      <ReactWebcam
        className={styles["cam"]}
        ref={ref}
        videoConstraints={{ width: 1280, height: 720 }}
        width={1280}
        height={720}
        screenshotFormat="image/jpeg"
      />
      <div className={styles["images"]}>
        {screenshots.map((screen) => (
          <img
            key={screen.createdAt}
            src={screen.path}
            width="200px"
            height="150px"
            alt={screen.createdAt}
          />
        ))}
      </div>
    </div>
  );
};

export default Webcam;
