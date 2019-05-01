package commands_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
	presenterfakes "github.com/pivotal-cf/om/presenters/fakes"
)

var _ = Describe("DiagnosticReport", func() {
	var (
		presenter   *presenterfakes.FormattedPresenter
		fakeService *fakes.DiagnosticReportService
		command     commands.DiagnosticReport
	)

	BeforeEach(func() {
		presenter = &presenterfakes.FormattedPresenter{}
		fakeService = &fakes.DiagnosticReportService{}
		command = commands.NewDiagnosticReport(presenter, fakeService)
	})

	Describe("Execute", func() {
		var diagnosticReport []api.DiagnosticReport

		BeforeEach(func() {
			diagnosticReport = []api.DiagnosticReport{}
		})

		It("displays the diagnostic report", func() {
			err := command.Execute([]string{})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeService.GetDiagnosticReportCallCount()).To(Equal(1))
			_ = diagnosticReport

			Expect(presenter.SetFormatArgsForCall(0)).To(Equal("json"))
			Expect(presenter.PresentDiagnosticReportCallCount()).To(Equal(1))
		})
	})

	Context("failure cases", func() {
		Context("when fetching the diagnostic report fails", func() {
			It("returns an error", func() {
				fakeService.GetDiagnosticReportReturns(api.DiagnosticReport{}, errors.New("beep boop"))

				err := command.Execute([]string{})
				Expect(err).To(MatchError("failed to retrieve diagnostic-report beep boop"))
			})
		})

		Context("when an unknown flag is passed", func() {
			It("returns an error", func() {
				err := command.Execute([]string{"--unknown-flag"})
				Expect(err).To(MatchError("could not parse diagnostic-report flags: flag provided but not defined: -unknown-flag"))
			})
		})
	})

	Describe("Usage", func() {
		It("returns usage information for the command", func() {
			command := commands.NewDiagnosticReport(nil, nil)
			Expect(command.Usage()).To(Equal(jhanda.Usage{
				Description:      "retrieve a diagnostic report with general information about the state of your Ops Manager.",
				ShortDescription: "reports current state of your Ops Manager",
				Flags:            command.Options,
			}))
		})
	})
})
