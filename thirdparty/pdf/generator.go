package pdf 

import (
  "context"
  "github.com/jung-kurt/gofpdf"
)

type PDFGenerator struct{} 

func (p *PDFGenerator) Generate(ctx context.Context, title, transcript, summary, output string) error {
  pdf := gofpdf.New("P", "mm", "A4", "")
  pdf.AddPage()

  pdf.SetFont("Arial", "B", 16)
  pdf.Cell(40, 10, title)

  pdf.SetFont("Arial", "", 12)
  pdf.Ln(12)
  pdf.MultiCell(0, 10, "Summary: \n" + summary, "", "", false)
  pdf.Ln(10)
  pdf.MultiCell(0, 10, "Transcript: \n"+ transcript, "", "", false)


  return pdf.OutputFileAndClose(output)
}
